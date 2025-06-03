/* Exercício 3 - Registro de clientes

O mesmo estudo do exercício anterior solicita uma funcionalidade para poder registrar novos dados de clientes. Os dados necessários são:

File
Name
ID
Phone number
Adress

Atividade 1: Antes de registrar um cliente, é necessário verificar se o cliente já existe. Para fazer isso, você precisa ler
 os dados de uma matriz. Caso ele se repita, você precisa tratar o erro adequadamente, como vimos até agora. Esse erro deve:
    - gerar um panic;
    - console iniciar a mensagem: “Error: client already exists”, e continuar com a execução do programa normalmente.

Atividade 2: Depois de tentar verificar se o cliente a ser registrado já existe, desenvolva uma função para validar se todos os
 dados a serem registrados para um cliente contêm um valor diferente de zero. Essa função deve retornar pelo menos dois valores.
 Um deles deverá ser do tipo erro, caso um valor zero seja inserido como parâmetro (lembre-se dos valores zero de cada tipo de dado,
 por exemplo: 0, “”, nil).

Atividade 3: Antes de encerrar a execução, mesmo que ocorram panics, as seguintes mensagens devem ser impressas no console: “End of
 execution” e “Several errors were detected at runtime”. Use o defer  para atender a esse requisito..

Requisitos gerais:

- Use a recover para recuperar o valor de qualquer pânico que possa ocorrer.
- Lembre-se de realizar as validações necessárias para cada retorno que possa conter um valor de erro.
- Gere um erro, personalizando-o de acordo com sua preferência usando uma das funções Go (execute também a validação relevante para
 o caso de erro retornado).
*/

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
