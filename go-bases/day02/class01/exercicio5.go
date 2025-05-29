package class01

import "fmt"

const (
	Dog       = "dog"
	Cat       = "cat"
	Tarantula = "tarantula"
	Hamster   = "hamster"
)

var foodPerAnimal = map[string]int{
	Dog:       10000,
	Cat:       5000,
	Tarantula: 250,
	Hamster:   150,
}

func Exercicio5() {
	fmt.Println("\nExercicio 5\n-----------------")

	animalDog, _ := Animal(Dog)
	animalCat, _ := Animal(Cat)
	animalTarantula, _ := Animal(Tarantula)
	animalHamster, _ := Animal(Hamster)
	animalPato, msg := Animal("pato")
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

func Animal(animalType string) (func(int) int, string) {
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
