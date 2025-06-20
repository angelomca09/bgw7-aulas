/* Exercício 4 - Quantos anos tem...

Um funcionário de uma empresa quer saber o nome e a idade de um de seus funcionários. De acordo com o mapa abaixo, ajude
 a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por outro lado, também é necessário:

- Saber quantos de seus funcionários têm mais de 21 anos.
- Adicionar um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Remover Pedro do mapa.
*/

package main

import "fmt"

func main() {
	fmt.Println("\nExercicio 4\n-----------------")

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Benjamin tem %d anos.\n", employees["Benjamin"])

	acimaDe21anos := 0

	// Calcula quantos funcionarios tem mais de 21 anos
	for _, idade := range employees {
		if idade > 21 {
			acimaDe21anos++
		}
	}

	var funcionarioText string
	if acimaDe21anos > 1 {
		funcionarioText = "funcionarios"
	} else {
		funcionarioText = "funcionario"
	}

	fmt.Printf("%d %s tem 21 anos.\n", acimaDe21anos, funcionarioText)

	// Adicionando um novo funcionário ao mapa
	employees["Federico"] = 25

	// Remover o funcionário "Pedro" do mapa corretamente
	delete(employees, "Pedro")

	fmt.Println(employees)
}
