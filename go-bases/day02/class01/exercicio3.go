package class01

import "fmt"

func Exercicio3() {
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
	fmt.Printf("Valor a receber: R$ %.2f\n", CalculaSalario(categoria, minutos))
}

func CalculaSalario(categoria string, minutos int) float64 {
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
