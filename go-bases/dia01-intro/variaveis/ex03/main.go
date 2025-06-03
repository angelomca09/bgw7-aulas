/* Exercício 3 - Declaração de variáveis

Um professor de programação está corrigindo os exames de seus alunos na disciplina Programação I para dar a eles os
 retornos correspondentes. Um dos itens do exame é declarar diferentes variáveis.

Preciso de ajuda para:

- Detecte quais dessas variáveis declaradas pelo aluno estão corretas.
- Corrija as incorretas.

var 1firstName string
var lastName string
var int age
1lastName := 6
var driver_license = true
var person height int
childsNumber := 2
*/

package main

import "fmt"

func main() {
	fmt.Println("Exercício 3\n-----------------")

	var firstName string

	var lastName string

	var age int

	var driver_license = true

	var person string

	var height int

	childsNumber := 2

	fmt.Println("firstName:", firstName)
	fmt.Println("lastName:", lastName)
	fmt.Println("age:", age)
	fmt.Println("driver_license:", driver_license)
	fmt.Println("person:", person)
	fmt.Println("height:", height)
	fmt.Println("childsNumber:", childsNumber)
}
