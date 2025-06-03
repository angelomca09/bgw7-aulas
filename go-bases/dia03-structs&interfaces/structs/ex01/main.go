/* Exercício 1

Crie um programa que atenda aos seguintes pontos:

- Ter uma estrutura chamada Product com os campos ID, Name, Price, Description e Category.
- Ter uma fatia global de Produto chamada Produtos instanciada com valores. 2 métodos associados à estrutura Produto: Save(), GetAll().
    O método Save() deve pegar a fatia de Products e adicionar o produto a partir do qual o método é chamado.
    O método GetAll() imprime todos os produtos salvos na fatia Products.
- Uma função getById() para a qual um INT deve ser passado como parâmetro e retorna o produto correspondente ao parâmetro passado.
- Execute pelo menos uma vez cada método e função definidos em main().
*/

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
