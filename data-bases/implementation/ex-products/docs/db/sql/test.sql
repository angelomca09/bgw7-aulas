DROP DATABASE IF EXISTS products_test_db;
CREATE DATABASE products_test_db;
USE products_test_db;

CREATE TABLE products (
    id INT DEFAULT NULL,
    name    varchar(50) DEFAULT NULL,
    quantity    int DEFAULT NULL,
    code_value  varchar(50) DEFAULT NULL,
    is_published    varchar(50) DEFAULT NULL,
    expiration  date DEFAULT NULL,
    price   decimal(5,2) DEFAULT NULL,
    id_warehouse INT DEFAULT NULL
);

CREATE TABLE warehouses (
  id    int DEFAULT NULL,
  name  varchar(255) DEFAULT NULL,
  address   varchar(150) DEFAULT NULL,
  telephone varchar(150) DEFAULT NULL,
  capacity  int DEFAULT NULL
);

INSERT INTO warehouses (id, name, address, telephone, capacity) VALUES
(1, 'Main Warehouse', '221 Baker Street', "4555666", 100),
(2, 'Secondary Warehouse', '334 Baker Street', "7888999", 150),
(3, 'Third Warehouse', '334 Baker Street', "7888999", 150),
(4, 'Fourth Warehouse', '334 Baker Street', "7888999", 150),
(5, 'Fifth Warehouse', '334 Baker Street', "7888999", 150);

INSERT INTO products (id, name, quantity, code_value, is_published, expiration, price, id_warehouse) VALUES
(1, 'product 1', 10, 'code 1', 'true', '2023-01-01', 100.00, 1),
(2, 'product 2', 20, 'code 2', 'true', '2023-01-01', 200.00, 2),
(3, 'product 3', 30, 'code 3', 'true', '2023-01-01', 300.00, 3),
(4, 'product 4', 40, 'code 4', 'true', '2023-01-01', 400.00, 4),
(5, 'product 5', 50, 'code 5', 'true', '2023-01-01', 500.00, 5);
