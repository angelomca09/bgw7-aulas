package class01

import "fmt"

func Exercicio02() {
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
