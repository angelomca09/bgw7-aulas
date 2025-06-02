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
