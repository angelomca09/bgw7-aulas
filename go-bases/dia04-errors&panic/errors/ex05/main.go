/* Exercício 5 -  Impostos sobre o salário #5

Vamos tornar nosso programa um pouco mais complexo e útil.

Desenvolver as funções necessárias para permitir que a empresa faça cálculos:
- Salário mensal de um trabalhador de acordo com o número de horas trabalhadas. A função deve receber as horas trabalhadas no mês
    e o valor da hora como argumento.
- A função deve retornar mais de um valor (salário calculado e erro). Caso o salário mensal seja igual ou superior a US$ 150.000,
    será deduzido um imposto de 10%. Caso o número de horas mensais inserido seja inferior a 80 ou um número negativo, a função
    deve retornar um erro. Ela indicará “Error: the worker cannot have worked less than 80 hours per month”.
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	salario, _ := calculaSalario(100, 1500)
	fmt.Println("Salario calculado para 101 horas trabalhadas e 1500 valor hora: ", salario)

	_, err := calculaSalario(79, 1500)
	if err != nil {
		fmt.Println(err)
	}
}

func calculaSalario(horasTrabalhadas, valorHora float64) (salario float64, err error) {
	if horasTrabalhadas < 80 {
		err = errors.New("Error: the worker cannot have worked less than 80 hours per month")
		return
	}
	salario = horasTrabalhadas * valorHora
	if salario >= 150000 {
		salario -= salario * 0.1
	}
	return
}
