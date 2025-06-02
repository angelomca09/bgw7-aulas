package main

import "fmt"

func main() {
	fmt.Println("\nExercicio 1\n-----------------")

	var palavra string
	fmt.Println("Digite a palavra: ")
	fmt.Scanln(&palavra)

	fmt.Printf("A palavra tem %d caracteres.\n", len(palavra))

	for _, char := range palavra {
		fmt.Printf("%c\n", char)
	}
}
