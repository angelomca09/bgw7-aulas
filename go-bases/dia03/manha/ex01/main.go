package main

import (
	"fmt"
)

var Produtos = []Product{
	{ID: 1, Name: "Mouse", Price: 20.90, Description: "Mouse Bluetooth", Category: "Periféricos"},
	{ID: 2, Name: "Monitor", Price: 1500.00, Description: "Monitor 4K", Category: "Eletrônicos"},
	{ID: 3, Name: "Camiseta", Price: 49.90, Description: "Camiseta de algodão", Category: "Vestuário"},
}

func main() {
	fmt.Println("Exercicio 01\n------------")

	p := Product{ID: 4, Name: "Teclado", Price: 89.90, Description: "Teclado mecânico", Category: "Periféricos"}
	p.Save()
	p.GetAll()
	fmt.Println("Produto com ID 2:", p.getById(2))
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Produtos = append(Produtos, p)
}

func (p Product) GetAll() {
	fmt.Println("Lista de Produtos: ")
	for _, product := range Produtos {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
	fmt.Println("")
}

func (p Product) getById(id int) Product {
	for _, product := range Produtos {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}
