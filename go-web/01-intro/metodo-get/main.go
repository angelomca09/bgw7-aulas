package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {
	router := chi.NewRouter()

	router.Get("/ping", PingHandler)
	router.Route("/products", func(r chi.Router) {
		r.Get("/", GetProductsHandler)            // Exercicio 2: Rota /products
		r.Get("/{id}", GetProductsByIDHandler)    // Exercicio 2: Rota /products/:id
		r.Get("/search", GetProuctsSearchHandler) // Exercicio 2: Rota /products/search
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic("Failed to start server: " + err.Error())
	}

}

// Exercicio 1: Receber os dados do JSON
func getProducts() []Product {
	var products []Product

	content, err := os.ReadFile("products.json")
	if err != nil {
		panic("Failed to read products.json: " + err.Error())
	}

	err = json.Unmarshal(content, &products)
	if err != nil {
		panic("Failed to unmarshal products.json: " + err.Error())
	}

	return products
}

// Exercicio 2: Criando um Servidor

// Crie uma rota /ping que deve responder com uma string contendo pong com status 200 OK.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// Crie uma rota /products que retorne a lista de todos os produtos na fatia.
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := getProducts()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	//w.Write(json.Marshal(products))
}

// Crie uma rota /products/:id que retorne um produto por seu ID.
func GetProductsByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	products := getProducts()

	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Crie uma rota /products/search que nos permita pesquisar por parâmetro os produtos cujo preço seja maior que um valor priceGt.
func GetProuctsSearchHandler(w http.ResponseWriter, r *http.Request) {
	priceGt := r.URL.Query().Get("priceGt")
	if priceGt == "" {
		http.Error(w, "priceGt query parameter is required", http.StatusBadRequest)
		return
	}
	// Convert priceGt to float64
	var priceGtFloat float64
	priceGtFloat, err := strconv.ParseFloat(priceGt, 64)
	if err != nil {
		http.Error(w, "Invalid priceGt value", http.StatusBadRequest)
		return
	}

	products := getProducts()
	var filteredProducts []Product

	for _, product := range products {
		if product.Price > priceGtFloat {
			filteredProducts = append(filteredProducts, product)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filteredProducts)
}
