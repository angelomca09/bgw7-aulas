-- Ex 1
-- Exibir o título e o nome do gênero de todas as séries.
SELECT se.title, ge.name
FROM series se INNER JOIN genres ge
	ON se.genre_id = ge.id

-- Ex 2
-- Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.
SELECT ep.title, ac.first_name, ac.last_name
FROM actors ac
	INNER JOIN actor_episode ac_ep ON ac.id = ac_ep.actor_id
	INNER JOIN episodes ep ON ac_ep.episode_id = ep.id

-- Ex 3
-- Mostre o título de todas as séries e o número total de temporadas de cada série.
SELECT se.title, COUNT(sea.serie_id)
FROM series se
	INNER JOIN seasons sea ON se.id = sea.serie_id
GROUP BY se.title

-- Ex 4
-- Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.
SELECT ge.name, count(mo.genre_id) as count
FROM genres ge
	INNER JOIN movies mo ON ge.id = mo.genre_id
GROUP BY ge.name
HAVING count > 3

-- Ex 5
-- Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.
SELECT DISTINCT ac.first_name, ac.last_name
FROM actors ac
	INNER JOIN actor_movie am ON ac.id = am.actor_id
    INNER JOIN movies mo ON mo.id = am.movie_id
WHERE mo.title LIKE 'La Guerra de las galaxias%'

