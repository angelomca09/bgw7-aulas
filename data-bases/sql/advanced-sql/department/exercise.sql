-- Ex 1
-- Selecciona o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
SELECT em.first_name, em.role, dp.locality
FROM employees em
	INNER JOIN departments dp ON em.dept_id = dp.id

-- Ex 2
-- Mostra os departamentos com mais de cinco empregados.
SELECT dp.name
FROM employees em
	INNER JOIN departments dp ON em.dept_id = dp.id
GROUP BY em.dept_id
HAVING count(*) > 5

-- Ex 3
-- Mostra o nome, o salário e o nome do departamento dos empregados que têm a mesma posição que "Mito Barchuk".
SELECT em.first_name, em.salary, dp.name
FROM employees em
	INNER JOIN departments dp ON em.dept_id = dp.id
WHERE em.role IN
	(SELECT em.role
		FROM employees em
			WHERE em.first_name = 'Mito' AND em.last_name = 'Barchuk')

-- Ex 4
-- Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.
SELECT em.*
FROM employees em
	INNER JOIN departments dp ON em.dept_id = dp.id
WHERE dp.name LIKE 'Contabilidade'

-- Ex 5
-- Mostra o nome do empregado com o salário mais baixo.
SELECT em.*
FROM employees em
WHERE em.salary IN (SELECT MIN(salary) FROM employees)

-- Ex 6
-- Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
SELECT em.*
FROM employees em
WHERE em.salary IN (SELECT MAX(em.salary) FROM employees em
	INNER JOIN departments dp ON em.dept_id = dp.id
WHERE dp.name LIKE 'Vendas');