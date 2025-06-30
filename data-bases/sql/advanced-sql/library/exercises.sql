-- Ex 1
-- Enumera os dados dos autores.
SELECT * FROM writers

-- Ex 2
-- Indica o nome e a idade dos alunos
SELECT first_name, age FROM students

-- Ex 3
-- Que alunos pertencem ao curso de informática?
SELECT first_name, last_name FROM students WHERE carreer = 'Informática'

-- Ex 4
-- Quais são os autores de nacionalidade francesa ou italiana?
SELECT name, nationality FROM writers WHERE nationality IN ('French', 'Italian')

-- Ex 5
-- Quais os livros que não são da área da Internet?
SELECT title, area FROM books WHERE area != 'Internet'

-- Ex 6
-- Enumera os livros publicados pela Salamandra.
SELECT title FROM books WHERE editorial = 'Salamandra'

-- Ex 7
-- Enumera os nomes dos alunos cuja idade é superior à média.
SELECT first_name, last_name FROM students WHERE age > (SELECT AVG(age) FROM students)

-- Ex 8
-- Enumera os nomes dos alunos cujo apelido começa com a letra G.
SELECT first_name, last_name FROM students WHERE last_name LIKE 'G%'

-- Ex 9
-- Faz uma lista dos autores do livro "O Universo: Guia de Viagem". (Apenas os nomes devem ser indicados).
SELECT w.name FROM writers w
    INNER JOIN books_writers bw ON w.id = bw.writer_id
    INNER JOIN books b ON b.id = bw.book_id
WHERE b.title = 'O Universo: Guia de Viagem'

-- Ex 10
-- Que livros emprestaste ao leitor "Filippo Galli"?
SELECT DISTINCT b.title FROM books b
    INNER JOIN rents r ON b.id = r.book_id
    INNER JOIN students s on s.id = r.reader_id
WHERE s.first_name = "Filippo" AND s.last_name = "Galli"

-- Ex 11
-- Indica o nome do aluno mais novo.
SELECT first_name, last_name FROM students
WHERE age IN (SELECT MIN(age) FROM students)

-- Ex 12
-- Enumera os nomes dos alunos a quem foram emprestados livros da Base de Dados.
SELECT first_name, last_name FROM students s
    INNER JOIN rents r ON s.id = r.reader_id
    INNER JOIN books b ON b.id = r.book_id
WHERE b.area = 'Base de Dados'

-- Ex 13
-- Enumera os livros que pertencem à autora J.K. Rowling.
SELECT b.title FROM books b
    INNER JOIN books_writers bw ON b.id = bw.book_id
    INNER JOIN writers w ON w.id = bw.writer_id
WHERE w.name = 'J.K. Rowling'

-- Ex 14
-- Enumera os títulos dos livros que deviam ser devolvidos em 16/07/2021.
SELECT b.title FROM books b
    INNER JOIN rents r ON b.id = r.book_id
WHERE r.return_date = '2021-07-16'