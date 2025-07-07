package repository

const (
	GetOneProductQuery = "SELECT id, name, quantity, code_value, is_published, expiration, price FROM products WHERE id = ?"
	StoreProductQuery  = "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?)"
	UpdateProductQuery = "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?"
	DeleteProductQuery = "DELETE FROM products WHERE id = ?"
)
