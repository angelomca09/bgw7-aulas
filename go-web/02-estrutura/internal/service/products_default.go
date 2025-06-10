package service

import (
	"estrutura-api/internal"
	"estrutura-api/internal/utils"
	"fmt"
)

type ProductsService struct {
	repository internal.ProductRepository
}

func NewProductService(repo internal.ProductRepository) internal.ProductService {
	service := &ProductsService{}

	service.repository = repo

	return service
}

func (p *ProductsService) GetAll() ([]*internal.Product, error) {
	products, err := p.repository.GetAll()
	return products, err
}

func (p *ProductsService) GetByID(id int) (*internal.Product, error) {
	product, err := p.repository.GetByID(id)
	return product, err
}

func (p *ProductsService) Create(product *internal.ProductAttributes) (*internal.Product, error) {

	if err := utils.ValidateEmptyProduct(product); err != nil {
		return nil, err
	}

	prs, _ := p.GetAll()
	if err := utils.ValidateCodeValue(product.CodeValue, prs); err != nil {
		return nil, err
	}

	if err := utils.ValidateExpiration(product.Expiration); err != nil {
		return nil, err
	}

	return p.repository.Create(product)
}

func (p *ProductsService) Update(id int, product *internal.Product) (*internal.Product, error) {
	originalProduct, err := p.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product.CodeValue != originalProduct.CodeValue {
		prs, _ := p.GetAll()
		if err := utils.ValidateCodeValue(product.CodeValue, prs); err != nil {
			return nil, err
		}
	}

	if err := utils.ValidateExpiration(product.Expiration); err != nil {
		return nil, err
	}

	return p.repository.Update(id, product)
}

func (p *ProductsService) UpdateOrCreate(product *internal.Product) (*internal.Product, error) {

	if err := utils.ValidateEmptyProduct(&product.ProductAttributes); err != nil {
		return nil, err
	}

	originalProduct, _ := p.repository.GetByID(product.ID)

	prs, _ := p.GetAll()
	if err := utils.ValidateCodeValue(product.CodeValue, prs); err != nil {
		if originalProduct == nil || product.CodeValue != originalProduct.CodeValue {
			return nil, err
		}
	}

	if err := utils.ValidateExpiration(product.Expiration); err != nil {
		return nil, err
	}

	return p.repository.UpdateOrCreate(product)
}

func (p *ProductsService) Delete(id int) error {
	err := p.repository.Delete(id)
	return err
}

func (p *ProductsService) GetTotalPriceByList(idList []int) ([]*internal.Product, float64, error) {
	totalPrice := 0.0
	askedList, err := p.repository.GetProductsByList(idList)
	if err != nil {
		return make([]*internal.Product, 0), 0.0, err
	}

	qntPerId := make(map[int]int, 0)
	for _, askedPr := range askedList {
		qntPerId[askedPr.ID] += 1
		if askedPr.Quantity < qntPerId[askedPr.ID] {
			return make([]*internal.Product, 0), 0.0, fmt.Errorf("Can not add more than there is in stock")
		}
		totalPrice += askedPr.Price
	}

	taxes := 1.21 // taxes for less than 10 products - 21%
	qntProd := len(askedList)
	switch {
	case qntProd > 10 && qntProd <= 20: // from 10 to 20 products - 17%
		taxes = 1.17
	case qntProd > 20: // more than 20 - 15%
		taxes = 1.15
	}

	totalPrice *= taxes

	return askedList, totalPrice, nil
}
