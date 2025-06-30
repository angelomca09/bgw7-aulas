-- Configura Database
DROP DATABASE IF EXISTS empresa_internet;
CREATE DATABASE empresa_internet;
USE empresa_internet;

-- Configura Tables
DROP TABLE IF EXISTS plans;
CREATE TABLE plans (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `speed` int  NOT NULL,
    `price` float  NOT NULL,
    `discount` float  NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `first_name` varchar(100)  NOT NULL,
    `last_name` varchar(100)  NOT NULL,
    `birthday` date NOT NULL,
    `province` varchar(100)  NOT NULL,
    `city` varchar(100)  NOT NULL,
    `plan_id` int unsigned,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`plan_id`) REFERENCES `plans`(`id`)
);

-- Dump
INSERT INTO plans (`id`, `speed`, `price`, `discount`) VALUES
(1, 100, 58.90, 0.0),
(2, 200, 98.90, 10.0),
(3, 350, 188.90, 20.0),
(4, 500, 398.90, 25.0),
(5, 1000, 798.90, 35.0);

INSERT INTO customers (`id`, `first_name`, `last_name`, `birthday`, `province`, `city`, `plan_id`) VALUES
(1, 'John', 'Doe', '1970-04-10', 'SP', 'Sao Paulo', 1),
(2, 'Marry', 'Doe', '1975-07-19', 'SP', 'Sorocaba', 3),
(3, 'Lewis', 'Doe', '1981-10-02', 'SC', 'Florianopolis', 3);

