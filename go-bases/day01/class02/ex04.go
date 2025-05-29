package class02

import "fmt"

func Exercicio4() {
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
