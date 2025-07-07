package internal

import "errors"

var (
	// ErrRepositoryWarehouseNotFound is returned when a warehouse is not found.
	ErrRepositoryWarehouseNotFound = errors.New("repository: warehouse not found")
)

// RepositoryWarehouse is an interface that contains the methods for a warehouse repository
type RepositoryWarehouse interface {
	// FindAll returns all warehouses
	FindAll() (p []Warehouse, err error)
	// FindById returns a warehouse by its id
	FindById(id int) (p Warehouse, err error)
	// Save saves a warehouse
	Save(p *Warehouse) (err error)
}
