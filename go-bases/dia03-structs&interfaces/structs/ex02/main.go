/* Exercício 2 - Employee

Uma empresa precisa fazer um bom gerenciamento de seus funcionários. Para isso, criaremos um pequeno programa que nos
ajudará a gerenciar corretamente esses funcionários. Os objetivos são:

- Criar uma estrutura Person com os campos ID, Name, DateOfBirth.
- Criar uma estrutura Employee com os campos: ID, Position e uma composição com a estrutura Person.
- Criar um método para a estrutura Employee chamado PrintEmployee(), que imprimirá os campos de um funcionário.
- Instanciar na função main() uma Person e um Employee carregando seus respectivos campos e, finalmente, executar o método PrintEmployee().

Se você conseguir realizar esse pequeno programa, poderá ajudar a empresa a resolver o problema de gerenciamento dos funcionários.
*/

package main

import "fmt"

func main() {
	fmt.Println("Exercicio 02\n------------")

	p := Person{ID: 1, Name: "John Doe", DateOfBirth: "1999-01-01"}
	e := Employee{ID: 1, Position: "Software Engineer", Person: p}
	e.PrintEmployee()
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	println("ID:", e.ID)
	println("Name:", e.Name)
	println("Date of Birth:", e.DateOfBirth)
	println("Position:", e.Position)
}
