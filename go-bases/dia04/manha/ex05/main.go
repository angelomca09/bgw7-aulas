package main

import (
	"errors"
	"fmt"
)

func main() {
	salario, _ := calculaSalario(100, 1500)
	fmt.Println("Salario calculado para 101 horas trabalhadas e 1500 valor hora: ", salario)

	_, err := calculaSalario(79, 1500)
	if err != nil {
		fmt.Println(err)
	}
}

func calculaSalario(horasTrabalhadas, valorHora float64) (salario float64, err error) {
	if horasTrabalhadas < 80 {
		err = errors.New("Error: the worker cannot have worked less than 80 hours per month")
		return
	}
	salario = horasTrabalhadas * valorHora
	if salario >= 150000 {
		salario -= salario * 0.1
	}
	return
}
