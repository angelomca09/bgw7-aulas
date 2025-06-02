package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("\nExercicio 2\n-----------------")

	fmt.Println("Calcule a média do aluno.\nDigite uma nota por vez (insira uma nota vazia para sair):")
	var notas []int
	var nota string
	for {
		nota = ""
		fmt.Scanln(&nota)
		n, err := strconv.Atoi(nota)
		if err != nil {
			break
		}
		notas = append(notas, n)
	}
	if len(notas) == 0 {
		fmt.Println("Nenhuma nota foi informada.")
		return
	}

	media := calculaMedia(notas)

	fmt.Printf("A média das notas é: %.2f\n", media)
}

func calculaMedia(notas []int) (media float64) {
	var soma float64
	for _, n := range notas {
		soma += float64(n)
	}
	media = soma / float64(len(notas))
	return
}
