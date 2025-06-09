package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"estrutura-api/internal"
	"fmt"
	"os"
)

type ProductJSONRepository struct {
	filePath string
	products map[int]*internal.ProductAttributes
	lastId   int
}

func NewProductJSONRepository() internal.ProductRepository {
	repo := &ProductJSONRepository{
		"docs/db/products.json",
		make(map[int]*internal.ProductAttributes),
		0,
	}

	repo.initialize_Decoder()

	return repo
}

func (r *ProductJSONRepository) initialize() {

	file, err := os.Open(r.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		product := &internal.Product{}

		if err := json.Unmarshal([]byte(line), &product); err != nil {
			line = line[:len(line)-1] // Removing ',' from the end of the line
			json.Unmarshal([]byte(line), &product)
		}

		r.products[product.ID] = &product.ProductAttributes
		r.lastId = len(r.products)
	}
	if err := scanner.Err(); err != nil {
		return
	}
}

func (r *ProductJSONRepository) initialize_Decoder() {
	file, err := os.Open(r.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	products := make([]internal.Product, 0)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		panic(err)
	}

	for _, product := range products {
		r.products[product.ID] = &product.ProductAttributes
	}
	r.lastId = len(r.products)
}

func (r *ProductJSONRepository) GetAll() ([]*internal.Product, error) {
	products := make([]*internal.Product, 0, len(r.products))
	for id, product := range r.products {
		products = append(products, &internal.Product{ID: id, ProductAttributes: (*product)})
	}
	return products, nil
}

func (r *ProductJSONRepository) GetByID(id int) (*internal.Product, error) {
	product := r.products[id]
	if product == nil {
		return nil, errors.New("Id invalid")
	}
	return &internal.Product{ID: id, ProductAttributes: *product}, nil
}

func (r *ProductJSONRepository) Create(product *internal.ProductAttributes) (*internal.Product, error) {
	if product == nil {
		return nil, errors.New("Invalid product")
	}

	r.lastId++
	r.products[r.lastId] = product

	return &internal.Product{ID: r.lastId, ProductAttributes: *product}, nil
}

func (r *ProductJSONRepository) Update(id int, product *internal.Product) (*internal.Product, error) {
	if _, ok := r.products[id]; !ok {
		err := fmt.Errorf("%s: %d", "Product not found", id)
		return nil, err
	}

	r.products[id] = &product.ProductAttributes

	return product, nil
}

func (r *ProductJSONRepository) UpdateOrCreate(product *internal.Product) (*internal.Product, error) {
	attr := internal.ProductAttributes{
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}

	id := 0
	_, ok := r.products[product.ID]
	switch ok {
	case true:
		r.products[product.ID] = &attr
		id = product.ID
	default:
		r.lastId++
		id = r.lastId
		r.products[id] = &attr
	}
	return &internal.Product{ID: id, ProductAttributes: *r.products[id]}, nil
}

func (r *ProductJSONRepository) Delete(id int) error {
	product := r.products[id]

	if product == nil {
		return errors.New("Invalid Id")
	}

	delete(r.products, id)

	return nil
}
