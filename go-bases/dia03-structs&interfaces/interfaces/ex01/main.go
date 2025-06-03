/*
	Exercício 1 - Registro de alunos

Uma universidade precisa registrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada aluno, como segue:

Nome: [Primeiro nome do aluno]
Sobrenome: [Sobrenome do aluno]
ID: [ID do aluno]
Data: [Data de admissão do aluno]

Os valores entre colchetes devem ser substituídos pelos dados fornecidos pelos alunos.
Para isso, é necessário gerar uma estrutura Students com as variáveis Name, Surname, DNI, Date e que tenha um métodode detalhamento
*/
package main

import "fmt"

func main() {
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
