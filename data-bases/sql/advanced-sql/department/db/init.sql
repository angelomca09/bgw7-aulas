-- Configura Database
DROP DATABASE IF EXISTS department_db;
CREATE DATABASE department_db;
USE department_db;

-- Configura Tables

DROP TABLE IF EXISTS departments;
CREATE TABLE departments (
    `id` varchar(7) NOT NULL,
    `name` varchar(100)  NOT NULL,
    `locality` varchar(100)  NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS employees;
CREATE TABLE employees (
    `id` varchar(6) NOT NULL,
    `first_name` varchar(100)  NOT NULL,
    `last_name` varchar(100)  NOT NULL,
    `role` varchar(100) NOT NULL,
    `created_at` date NOT NULL,
    `salary` float NOT NULL,
    `comission` float NOT NULL,
    `dept_id` varchar(7),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`dept_id`) REFERENCES `departments`(`id`)
);

-- Dump
INSERT INTO departments (`id`, `name`, `locality`) VALUES
('D-000-1', 'Software', 'Los Tigres'),
('D-000-2', 'Sistemas', 'Guadalupe'),
('D-000-3', 'Contabilidade', 'La Roca'),
('D-000-4', 'Vendas', 'Plata');

INSERT INTO employees (`id`, `first_name`, `last_name`, `role`, `created_at`, `salary`, `comission`, `dept_id`) VALUES
('E-0001', 'César', 'Piñero', 'Vendedor',  '2018-05-12', 80000, 15000, 'D-000-4'),
('E-0002', 'Yosep', 'Kowaleski', 'Analista',  '2015-07-14', 140000, 0, 'D-000-2'),
('E-0003', 'Mariela', 'Barrios', 'Diretor',  '2014-06-05', 185000, 0, 'D-000-3'),
('E-0004', 'Jonathan', 'Aguilera', 'Vendedor',  '2015-06-03', 85000, 10000, 'D-000-4'),
('E-0005', 'Daniel', 'Brezezicki', 'Vendedor',  '2018-03-03', 83000, 10000, 'D-000-4'),
('E-0006', 'Mito', 'Barchuk', 'Presidente',  '2014-06-05', 190000, 0, 'D-000-3'),
('E-0007', 'Emilio', 'Galarza', 'Desenvolvedor',  '2014-08-02', 60000, 0, 'D-000-1');

