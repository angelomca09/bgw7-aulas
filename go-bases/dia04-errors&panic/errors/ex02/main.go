/* Exercício 2 - Impostos sobre o salário #2

Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int". Crie um erro personalizado
com uma estrutura que implemente "Error()" com a mensagem "Error: salary is less than 10000" e a inicie caso "salary" seja menor
ou igual a 10000. A validação deve ser feita com a função "Is()" dentro do "main".
*/

package main

import "errors"

func main() {
	salary := 1000

	var err *myError

	if salary < 10000 {
		err = &myError{}
	}

	if errors.Is(err, &myError{}) {
		panic(err)
	}
}

type myError struct{}

func (e *myError) Error() string {
	return "Error: salary is less than 10000"
}
