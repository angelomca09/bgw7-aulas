/* Exercício 1 - Letras em uma palavra

A Real Academia Espanhola quer saber quantas letras uma palavra tem e, em seguida, ter cada uma das letras
separadamente para soletrá-la. Para isso, eles terão de:
- Criar um aplicativo que receba uma variável com a palavra e imprima o número de letras que ela contém.
- Em seguida, imprima cada uma das letras.
*/

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
