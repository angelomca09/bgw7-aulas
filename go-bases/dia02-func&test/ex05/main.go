/* Exercício 5 - Calcular a quantidade de alimentos

Um abrigo de animais precisa calcular a quantidade de alimentos que precisa comprar para seus animais de estimação.
No momento, eles só têm tarântulas, hamsters, cães e gatos, mas esperam poder abrigar muito mais animais.
Eles têm os seguintes dados:

- Cão 10 kg de comida.
- Gato 5 kg de comida.
- Hamster 250 g de comida.
- Tarântula 150 g de comida.

É solicitado que você:

- Implemente uma função Animal que receba como parâmetro um valor de texto com o animal especificado e retorne uma função
 e uma mensagem (caso o animal não exista).
- Uma função para cada animal que calcule a quantidade de comida com base na quantidade do tipo de animal especificado.

Exemplo:

const (
   dog    = "dog"
   cat    = "cat"
)

...

animalDog, msg := animal(dog)
animalCat, msg := animal(cat)

...

var amount float64
amount += animalDog(10)
amount += animalCat(10)
*/

package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	tarantula = "tarantula"
	hamster   = "hamster"
)

var foodPerAnimal = map[string]int{
	dog:       10000,
	cat:       5000,
	tarantula: 250,
	hamster:   150,
}

func main() {
	fmt.Println("\nExercicio 5\n-----------------")

	animalDog, _ := animal(dog)
	animalCat, _ := animal(cat)
	animalTarantula, _ := animal(tarantula)
	animalHamster, _ := animal(hamster)
	animalPato, msg := animal("pato")
	if msg != "" {
		fmt.Println(msg)
	}

	var amount int

	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalTarantula(10)
	amount += animalHamster(10)
	amount += animalPato(10) // Retorno zerado por se tratar de um animal que não existe no mapa

	fmt.Printf("Total de ração necessária: %.2f Kg\n", float64(amount/1000))
}

func animal(animalType string) (func(int) int, string) {
	var msg string = ""
	var foodPerKg int = 0.0

	if foodPerAnimal[animalType] == 0 {
		msg = "Animal " + animalType + " não cadastrado."
	} else {
		foodPerKg = foodPerAnimal[animalType]
	}

	return func(animalQuantity int) int {
		return foodPerKg * animalQuantity
	}, msg
}
