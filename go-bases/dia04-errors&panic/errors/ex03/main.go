/* Exercício 3 - Impostos sobre o salário #3

Faça o mesmo que no exercício anterior, mas reformule o código de modo que, em vez de "Error()", seja implementado "errors.New()".
*/

package main

import "errors"

func main() {
	salary := 1000

	var err error

	if salary < 10000 {
		err = errors.New("Error: salary is less than 10000")
	}

	if errors.Is(err, err) {
		panic(err)
	}
}
