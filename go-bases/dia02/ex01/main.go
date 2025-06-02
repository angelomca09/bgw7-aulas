package main

import "fmt"

func main() {
	fmt.Println("\nExercicio 1\n-----------------")

	var salario float64
	fmt.Print("Digite o salário do funcionário: ")
	fmt.Scan(&salario)

	var imposto, porcentagem = calculaImposto(salario)

	fmt.Printf("O imposto a ser pago é de R$%.2f, correspondente a %.0f%% do salário.\n", imposto, porcentagem*100)

}

func calculaImposto(salario float64) (imposto, porcentagem float64) {
	porcentagem = 0.0
	if salario > 50000 {
		porcentagem += 0.17
	}
	if salario > 150000 {
		porcentagem += 0.10
	}
	return salario * porcentagem, porcentagem
}
