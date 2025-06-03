/*
	Exercício 3 - Cálculo de Salário

Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas trabalhadas por mês e na categoria.

- Categoria C, o salário deles é de US$ 1.000 por hora.
- Categoria B, o salário deles é de US$ 1.500 por hora, mais 20% do salário mensal.
- Categoria A, o salário deles é de US$ 3.000 por hora, mais 50% do salário mensal.

Você deve gerar uma função que receba como parâmetro o número de minutos trabalhados por mês, a categoria e retorne o salário.
*/
package main

import "fmt"

func main() {
	fmt.Println("\nExercicio 3\n-----------------")

	var categoria string
	var minutos int
	for {
		fmt.Println("Informe a Categoria (A, B ou C) e a quantidade de minutos trabalhados:")
		fmt.Scanln(&categoria, &minutos)

		switch categoria {
		case "A", "a":
			fallthrough
		case "B", "b":
			fallthrough
		case "C", "c":
			break
		default:
			fmt.Println("Categoria inválida. Tente novamente.")
			continue
		}

		if minutos < 0 {
			fmt.Println("Quantidade de minutos inválida. Deve ser um número positivo.")
			continue
		}

		break
	}
	fmt.Printf("Valor a receber: R$ %.2f\n", calculaSalario(categoria, minutos))
}

func calculaSalario(categoria string, minutos int) float64 {
	var valorHora float64
	var adicionalmensal float64 = 0.00

	switch categoria {
	case "A", "a":
		valorHora = 3000.00
		adicionalmensal = 0.5
	case "B", "b":
		valorHora = 1500.00
		adicionalmensal = 0.2
	case "C", "c":
		valorHora = 1000.00
	}

	valorTotal := (float64(minutos) / 60) * valorHora
	valorTotal += valorTotal * adicionalmensal
	return valorTotal
}
