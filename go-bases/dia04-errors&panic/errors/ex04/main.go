/* Exercício 4 - Impostos sobre o salário #4

Repeti o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de erro receba como parâmetro o valor
de "salary" indicando que ele não atinge o mínimo tributável (a mensagem exibida pelo console deve dizer:“Error: the minimum
taxable amount is 150,000 and the salary entered is: [salary]”, sendo  [salary] o valor do tipo int passado pelo parâmetro).
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 1000

	var err error

	if salary < 15000 {
		err = fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}

	if errors.Is(err, fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)) {
		panic(err)
	}
}
