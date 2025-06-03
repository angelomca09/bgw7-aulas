/* Exercício 1 - Impostos sobre o salário #1

Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".

Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: the salary entered does not reach
the taxable minimum" e acione-a caso "salary" seja menor que 150.000. Caso contrário, será necessário imprimir no console a
mensagem "Must pay tax".
*/

package main

import "fmt"

func main() {
	salary := 123456

	if salary < 150000 {
		err := &myError{}
		panic(err)
	} else {
		fmt.Println("Must pay taxes")
	}
}

type myError struct{}

func (e *myError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}
