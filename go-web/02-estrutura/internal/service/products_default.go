package service

import (
	"estrutura-api/internal"
	"estrutura-api/internal/utils"
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

	if err := p.validateProduct(product); err != nil {
		return nil, err
	}

	return p.repository.Create(product)
}

func (p *ProductsService) Update(id int, product *internal.Product) (*internal.Product, error) {
	if err := p.validateProduct(&product.ProductAttributes); err != nil {
		return nil, err
	}

	return p.repository.Update(id, product)
}

func (p *ProductsService) UpdateOrCreate(product *internal.Product) (*internal.Product, error) {
	if err := p.validateProduct(&product.ProductAttributes); err != nil {
		return nil, err
	}

	return p.repository.UpdateOrCreate(product)
}

func (p *ProductsService) Delete(id int) error {
	err := p.repository.Delete(id)
	return err
}

func (p *ProductsService) validateProduct(product *internal.ProductAttributes) error {
	prs, _ := p.GetAll()
	if err := utils.ValidateCodeValue(product.CodeValue, prs); err != nil {
		return err
	}

	if err := utils.ValidateExpiration(product.Expiration); err != nil {
		return err
	}
	return nil
}
