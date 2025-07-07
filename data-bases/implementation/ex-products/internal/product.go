package internal

import "time"

// ProductAttributes is a struct that contains the attributes of a product
type ProductAttributes struct {
	// Name is the name of the product
	Name string `json:"name"`
	// Quantity is the quantity of the product
	Quantity int `json:"quantity"`
	// CodeValue is the code value of the product
	CodeValue string `json:"code_value"`
	// IsPublished is the published status of the product
	IsPublished bool `json:"is_published"`
	// Expiration
	Expiration time.Time `json:"expiration"`
	// Price
	Price float64 `json:"price"`
	// Warehouse Id
	WarehouseId int `json:"warehouse_id"`
}

// Product is a struct that contains the attributes of a product
type Product struct {
	// Id is the unique identifier of the product
	Id int `json:"id"`
	// ProductAttributes is the attributes of the product
	ProductAttributes
}
