package class02

import "fmt"

func Exercicio2() {
	println("\nExercicio 2\n-----------------")

	var idade int
	var empregado string
	var anosEmpregado int
	var salario float64

	//Maior de 22 anos
	fmt.Println("Digite a idade: ")
	fmt.Scanln(&idade)

	if idade < 22 {
		fmt.Println("O Banco só concede empréstimo para maiores de 22 anos.")
		return
	}

	//Empregado
	for {
		fmt.Println("Está empregado? (s/n): ")
		fmt.Scanln(&empregado)

		switch empregado {
		case "s", "S":
			break
		case "n", "N":
			fmt.Println("O Banco só concede empréstimo para empregados.")
			return
		default:
			fmt.Println("Opção inválida. Digite 's' para sim ou 'n' para não.")
			continue
		}

		// Se o usuário digitou 's' ou 'S', sai do loop
		if empregado == "s" || empregado == "S" {
			break
		}
	}

	fmt.Println("Quantos anos está empregado? ")
	fmt.Scanln(&anosEmpregado)
	if anosEmpregado < 1 {
		fmt.Println("O Banco só concede empréstimo para empregados com mais de 1 ano de serviço.")
		return
	}

	//Salario
	fmt.Println("Qual seu salário?")
	fmt.Scanln(&salario)
	if salario < 100000 {
		fmt.Println("O Banco só concede empréstimo para salários acima de R$100.000,00.")
		return
	}

	fmt.Println("Empréstimo concedido com sucesso!")
}
