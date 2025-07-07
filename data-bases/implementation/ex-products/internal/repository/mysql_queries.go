package repository

const (
	GetOneProductQuery = "SELECT id, name, quantity, code_value, is_published, expiration, price, id_warehouse FROM products WHERE id = ?"
	StoreProductQuery  = "INSERT INTO products (name, quantity, code_value, is_published, expiration, price, id_warehouse) VALUES (?, ?, ?, ?, ?, ?, ?)"
	UpdateProductQuery = "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ?, id_warehouse = ? WHERE id = ?"
	DeleteProductQuery = "DELETE FROM products WHERE id = ?"
)
