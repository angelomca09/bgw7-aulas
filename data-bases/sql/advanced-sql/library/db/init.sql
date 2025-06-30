DROP DATABASE IF EXISTS books_db;
CREATE DATABASE books_db;
USE books_db;

DROP TABLE IF EXISTS writers;
CREATE TABLE writers (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `nationality` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS books;
CREATE TABLE books (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `editorial` varchar(100) NOT NULL,
    `area` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS students;
CREATE TABLE students (
    `id` int NOT NULL AUTO_INCREMENT,
    `first_name` varchar(100) NOT NULL,
    `last_name` varchar(100) NOT NULL,
    `email` varchar(100) NOT NULL,
    `carreer` varchar(100) NOT NULL,
    `age` int NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS books_writers;
CREATE TABLE books_writers (
    `book_id` int NOT NULL,
    `writer_id` int NOT NULL,
    PRIMARY KEY (`book_id`, `writer_id`),
    FOREIGN KEY (`book_id`) REFERENCES `books`(`id`),
    FOREIGN KEY (`writer_id`) REFERENCES `writers`(`id`)
);

DROP TABLE IF EXISTS rents;
CREATE TABLE rents (
    `id` int NOT NULL AUTO_INCREMENT,
    `reader_id` int NOT NULL,
    `book_id` int NOT NULL,
    `rent_date` date NOT NULL,
    `return_date` date NOT NULL,
    `returned` boolean NOT NULL DEFAULT false,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`reader_id`) REFERENCES `students`(`id`),
    FOREIGN KEY (`book_id`) REFERENCES `books`(`id`)
);

-- Dump
INSERT INTO writers (`id`, `name`, `nationality`) VALUES
(1, 'Gabriel García Márquez', 'Colombian'),
(2, 'J.K. Rowling', 'British'),
(3, 'George Orwell', 'British'),
(4, 'Jane Austen', 'British'),
(5, 'Mark Twain', 'American');

INSERT INTO books (`id`, `title`, `editorial`, `area`) VALUES
(1, 'Cien años de soledad', 'Editorial Sudamericana', 'Fiction'),
(2, 'Harry Potter and the Philosophers Stone', 'Bloomsbury', 'Fantasy'),
(3, '1984', 'Secker & Warburg', 'Dystopian'),
(4, 'Pride and Prejudice', 'T. Egerton', 'Romance'),
(5, 'The Adventures of Tom Sawyer', 'Chatto & Windus', 'Adventure');

INSERT INTO students (`id`, `first_name`, `last_name`, `email`, `carreer`, `age`) VALUES
(1, 'Alice', 'Johnson', 'alice.johson@email.com', 'Computer Science', 20),
(2, 'Bob', 'Smith', 'bob.smith@email.com', 'Literature', 22),
(3, 'Charlie', 'Brown', 'charlie.brown@email.com', 'History', 21),
(4, 'Diana', 'Prince', 'diana.prince@email.com', 'Physics', 23),
(5, 'Ethan', 'Hunt', 'ethan.hunt@email.com', 'Engineering', 24);

INSERT INTO books_writers (`book_id`, `writer_id`) VALUES
(1, 1), -- Cien años de soledad by Gabriel García Márquez
(2, 2), -- Harry Potter and the Philosophers Stone by J.K. Rowling
(3, 3), -- 1984 by George Orwell
(4, 4), -- Pride and Prejudice by Jane Austen
(5, 5); -- The Adventures of Tom Sawyer by Mark Twain

INSERT INTO rents (`id`, `reader_id`, `book_id`, `rent_date`, `return_date`, `returned`) VALUES
(1, 1, 1, '2023-01-10', '2023-01-20', true), -- Alice rented Cien años de soledad
(2, 2, 2, '2023-02-15', '2023-02-25', false), -- Bob rented Harry Potter and the Philosophers Stone
(3, 3, 3, '2023-03-05', '2023-03-15', true), -- Charlie rented 1984
(4, 4, 4, '2023-04-01', '2023-04-11', false), -- Diana rented Pride and Prejudice
(5, 5, 5, '2023-05-10', '2023-05-20', true); -- Ethan rented The Adventures of Tom Sawyer
