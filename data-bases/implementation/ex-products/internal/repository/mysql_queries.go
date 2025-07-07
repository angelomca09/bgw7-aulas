package repository

const (
	GetOneProductQuery = "SELECT id, name, quantity, code_value, is_published, expiration, price, id_warehouse FROM products WHERE id = ?"
	StoreProductQuery  = "INSERT INTO products (name, quantity, code_value, is_published, expiration, price, id_warehouse) VALUES (?, ?, ?, ?, ?, ?, ?)"
	UpdateProductQuery = "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ?, id_warehouse = ? WHERE id = ?"
	DeleteProductQuery = "DELETE FROM products WHERE id = ?"

	GetOneWarehouseQuery = "SELECT id, name, address, telephone, capacity FROM warehouses WHERE id = ?"
	StoreWarehouseQuery  = "INSERT INTO warehouses (name, address, telephone, capacity) VALUES (?, ?, ?, ?)"

	GetInnerJoinProductWarehouseQuery = "SELECT w.name, SUM(p.quantity) FROM products p INNER JOIN warehouses w ON p.id_warehouse = w.id"
)
