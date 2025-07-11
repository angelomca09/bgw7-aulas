-- Uso base de datos `storage`
-- USE `storage`;

-- Crear tabla warehouses donde se almacenan datos de los almacenes de products   
CREATE TABLE `warehouses` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `address` varchar(150) NOT NULL,
  `telephone` varchar(150) NOT NULL,
  `capacity` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Volcado de datos para la tabla `warehouses`
INSERT INTO `warehouses` (`id`, `name`, `address`, `telephone`, `capacity`) VALUES
(1, 'Main Warehouse', '221 Baker Street', "4555666", 100),
(2, 'Secondary Warehouse', '334 Baker Street', "7888999", 150);

-- Asignar la columna `id` 
ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`);

-- Modificar tabla warehouses para que el id sea autoincrementable
ALTER TABLE `warehouses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

-- Se agrega una columna para el id del warehouse a la tabla products
ALTER TABLE `products`
    ADD `id_warehouse` INT NULL AFTER `price`,
    ADD CONSTRAINT `products_warehouses_id_fk` FOREIGN KEY (`id_warehouse`) REFERENCES `warehouses` (`id`);

-- Designar el id_warehouse con el valor 1
UPDATE `products` SET `id_warehouse` = id % 2 + 1;