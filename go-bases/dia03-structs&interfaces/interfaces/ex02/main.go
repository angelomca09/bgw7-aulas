/*
	Exercício 2 - Produtos

Algumas lojas de comércio eletrônico precisam criar uma funcionalidade no Go para gerenciar produtos e retornar o valor do preço total.
A empresa tem três tipos de produtos: Pequeno, Médio e Grande (muitos outros são esperados).

E os custos adicionais são:

Pequeno: apenas o custo do produto
Médio: o preço do produto + 3% do produto + 3% de mantê-lo na loja
Grande: o preço do produto + 6% de mantê-lo na loja e, além disso, o custo de envio de US$ 2.500.

O custo de manter o produto em estoque na loja é uma porcentagem do preço do produto.

É necessária uma função factory que receba o tipo de produto e o preço e retorne uma interface Product que tenha o método Price.

Deve ser possível executar o método Price e fazer com que o método retorne o preço total com base no custo do produto e em quaisquer
custos adicionais.
*/
package main

import "fmt"

func main() {
	fmt.Println("Exercicio 02\n------------")

	// Criando produtos de diferentes tamanhos
	smallProduct := Factory("small", 100.0)
	mediumProduct := Factory("medium", 100.0)
	largeProduct := Factory("large", 100.0)

	// Exibindo os preços dos produtos
	fmt.Println("Preço do produto pequeno:", smallProduct.Price())
	fmt.Println("Preço do produto médio:", mediumProduct.Price())
	fmt.Println("Preço do produto grande:", largeProduct.Price())
}

type Product interface {
	Price() float64
}

type Products struct {
	BasePrice float64
	Product
}

type SmallProduct struct {
	Products
}

func (p *SmallProduct) Price() float64 {
	// Pequeno: apenas o custo do produto
	return p.BasePrice
}

type MediumProduct struct {
	Products
}

func (p *MediumProduct) Price() float64 {
	// Médio: o preço do produto + 3% do produto + 3% de mantê-lo na loja
	return p.BasePrice + p.BasePrice*0.03 + p.BasePrice*0.03
}

type LargeProduct struct {
	Products
}

func (p *LargeProduct) Price() float64 {
	// Grande: o preço do produto + 6% de mantê-lo na loja e, além disso, o custo de envio de US$ 2.500
	return p.BasePrice + p.BasePrice*0.06 + 2500
}

func Factory(productType string, basePrice float64) Product {
	switch productType {
	case "small":
		return &SmallProduct{
			Products: Products{
				BasePrice: basePrice,
			},
		}
	case "medium":
		return &MediumProduct{
			Products: Products{
				BasePrice: basePrice,
			},
		}
	case "large":
		return &LargeProduct{
			Products: Products{
				BasePrice: basePrice,
			},
		}
	default:
		return nil
	}
}
