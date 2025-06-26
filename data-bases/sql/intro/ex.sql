--Ex 1
--Mostrar todos os registros da tabela de filmes.
SELECT * FROM movies

--Ex 2
--Exibir o primeiro nome, o sobrenome e a classificação de todos os atores.
SELECT a.first_name, a.last_name, a.rating FROM actors as a

--Ex 3
--Exibir o título de todas as séries e usar aliases para que o nome da tabela e o campo estejam em inglês.
SELECT s.title as title FROM series as s

--Ex 4
--Exibir o primeiro nome e o sobrenome dos atores cuja classificação seja superior a 7,5.
SELECT a.first_name, a.last_name FROM actors as a WHERE rating > 7.5

--Ex 5
--Exibir o título do filme, a classificação e os prêmios dos filmes com classificação superior a 7,5 e com mais de dois prêmios.
SELECT m.title, m.rating FROM movies as m WHERE m.rating > 7.5 AND m.awards > 2

--Ex 6
--Exibir o título do filme e a classificação ordenados por classificação em ordem crescente.
SELECT m.title, m.rating FROM movies as m ORDER BY m.rating ASC

--Ex 7
--Exibir os títulos dos filmes e a classificação ordenados por classificação.
SELECT m.title, m.rating FROM movies as m ORDER BY m.rating

--Ex 8
--Exibir os títulos dos três primeiros filmes no banco de dados.
SELECT m.title FROM movies AS m LIMIT 3

--Ex 9
--Exibir os 5 filmes mais bem classificados.
SELECT m.title FROM movies AS m ORDER BY m.rating DESC LIMIT 5

--Ex 10
--Listar os 10 primeiros atores.
SELECT * FROM actors AS a LIMIT 10

--Ex 11
--Exibir o título e a classificação de todos os filmes cujo título é Toy Story.
SELECT m.title, m.rating FROM movies as m WHERE m.title = 'Toy Story'

--Ex 12
--Exibir todos os atores cujos nomes começam com Sam.
SELECT * FROM actors AS a WHERE a.first_name LIKE 'Sam%'

--Ex 13
--Exibir o título dos filmes lançados entre 2004 e 2008.
SELECT m.title FROM movies AS m WHERE m.release_date BETWEEN '2004-01-01' AND '2008-01-01'

--Ex 14
--Exibir o título dos filmes com classificação superior a 3, com mais de 1 prêmio e com data de lançamento entre 1988 e 2009.
SELECT m.title FROM movies AS m WHERE m.release_date BETWEEN '2004-01-01' AND '2008-01-01'

--Ex 15
--Ordenar os resultados por classificação.
SELECT m.title, m.rating FROM movies AS m WHERE m.release_date BETWEEN '1988-01-01' AND '2009-01-01' ORDER BY m.rating