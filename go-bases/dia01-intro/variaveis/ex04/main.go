/* Exercício 4 - Tipos de dados

Um estudante de programação tentou fazer declarações de variáveis com seus respectivos tipos em Go, mas teve várias
 dúvidas ao fazer isso. A partir disso, ele nos forneceu seu código e pediu a ajuda de um desenvolvedor experiente que pudesse:

Verifique seu código e faça as correções necessárias.

var lastName string = "Smith"
var age int = "35"
boolean := "false"
var salary string = 45857.90
var firstName string = "Mary"
*/

package main

import "fmt"

func main() {
	fmt.Println("Exercício 4\n-----------------")

	var lastName string = "Smith"

	var age int = 35

	boolean := false

	var salary float64 = 45857.90

	var firstName string = "Mary"

	fmt.Println("lastName:", lastName)
	fmt.Println("age:", age)
	fmt.Println("boolean:", boolean)
	fmt.Println("salary:", salary)
	fmt.Println("firstName:", firstName)
}
