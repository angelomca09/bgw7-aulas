package internal

type ProductService interface {
	GetAll() ([]*Product, error)
	GetByID(id int) (*Product, error)
	Create(product *ProductAttributes) (*Product, error)
	Update(id int, product *Product) (*Product, error)
	UpdateOrCreate(product *Product) (*Product, error)
	Delete(id int) error
	GetTotalPriceByList(id []int) ([]*Product, float64, error)
}
