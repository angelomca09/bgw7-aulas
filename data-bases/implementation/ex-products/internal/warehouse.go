package internal

// WarehouseAttributes is a struct that contains the attributes of a warehouse
type WarehouseAttributes struct {
	// Name is the name of the warehouse
	Name string `json:"name"`
	// Address is the address of the warehouse
	Address string `json:"address"`
	// Telephone is the telephone of the warehouse
	Telephone string `json:"telephone"`
	// Capacity is the capacity of the warehouse
	Capacity int `json:"capacity"`
}

// Warehouse is a struct that contains the attributes of a warehouse
type Warehouse struct {
	// Id is the unique identifier of the warehouse
	Id int `json:"id"`
	// WarehouseAttributes is the attributes of the warehouse
	WarehouseAttributes
}
