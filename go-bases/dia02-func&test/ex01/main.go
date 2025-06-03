/* Exercício 1 - Impostos sobre o salário

Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário.
Para cumprir o objetivo, é necessário criar uma função que retorne o imposto de um salário.

Levando em conta que se a pessoa ganha mais de US$ 50.000, 17% do salário será deduzido e se a
 pessoa ganha mais de US$ 150.000, 10% também será deduzido (27% no total).
*/

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
