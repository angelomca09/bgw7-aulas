package class02

import "fmt"

func Exercicio01() {
	fmt.Println("Exercicio 01\n------------")

	s := Student{
		ID:            1,
		Name:          "John",
		Surname:       "Doe",
		AdmissionDate: "2023-01-01",
	}

	s.ShowInfo()
}

type Student struct {
	ID            int
	Name          string
	Surname       string
	AdmissionDate string
}

func (s *Student) ShowInfo() {
	fmt.Println("Nome:", s.Name)
	fmt.Println("Sobrenome:", s.Surname)
	fmt.Println("ID:", s.ID)
	fmt.Println("Data:", s.AdmissionDate)
}
