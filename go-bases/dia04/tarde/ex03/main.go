package main

import (
	"errors"
	"fmt"
)

var clientes Clientes = []Cliente{
	{File: "johndoe1.txt", Name: "John Doe", ID: 1, PhoneNumber: "123-456-7890", Adress: "123 Elm St"},
	{File: "janesmith2.txt", Name: "Jane Smith", ID: 2, PhoneNumber: "987-654-3210", Adress: "456 Oak St"},
	{File: "alicejohnson3.txt", Name: "Alice Johnson", ID: 3, PhoneNumber: "555-123-4567", Adress: "789 Pine St"},
	{File: "bobbrown4.txt", Name: "Bob Brown", ID: 4, PhoneNumber: "444-987-6543", Adress: "321 Maple St"},
}

var errqtd int = 0

func main() {
	defer func() {
		fmt.Println("End of execution")
	}()
	newCliente := Cliente{
		File:        "",
		Name:        "",
		ID:          1,
		PhoneNumber: "",
		Adress:      "",
	}

	// Verificar se o cliente existe
	clientes.hasCliente(newCliente.ID)

	// Validar se os dados do cliente não estão vazios
	validateCliente(newCliente)

	if errqtd > 0 {
		defer func() {
			fmt.Println("Several errors were deteced at runtime")
		}()
	}
}

type Cliente struct {
	File        string
	Name        string
	ID          int
	PhoneNumber string
	Adress      string
}

type Clientes []Cliente

func (cl *Clientes) hasCliente(id int) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Method hasCliente recovered from panic")
			fmt.Println("Recovered from panic:", err)
			errqtd++
		}
	}()

	if id <= 0 {
		return false, fmt.Errorf("invalid ID: %d", id)
	}

	for _, cliente := range *cl {
		if cliente.ID == id {
			panic(errors.New("Cliente already exists with ID: " + fmt.Sprint(id)))
		}
	}
	return false, nil
}

func validateCliente(cliente Cliente) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Method validateCliente recovered from panic")
			fmt.Println("Recovered from panic:", err)
			errqtd++
		}
	}()

	if cliente.Name == "" || cliente.PhoneNumber == "" || cliente.Adress == "" ||
		cliente.File == "" || cliente.ID == 0 {
		panic(errors.New("cliente data cannot be empty"))
	}
}
