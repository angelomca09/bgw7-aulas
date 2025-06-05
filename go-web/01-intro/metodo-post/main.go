package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	ct := NewController()
	rt := chi.NewRouter()

	rt.Route("/products", func(rt chi.Router) {
		rt.Post("/", ct.CreateProducts())
		rt.Get("/{id}", ct.GetProductById())
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type RequestBodyProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseBodyProduct struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
	Error   bool     `json:"error"`
}

type Controller interface {
	CreateProducts() http.HandlerFunc
	GetProductById() http.HandlerFunc
}

type ControllerProducts struct {
	storage map[int]*Product
}

func NewController() Controller {
	return &ControllerProducts{storage: make(map[int]*Product, 0)}
}

func (c *ControllerProducts) ValidateCodeValue(cv string) error {
	if cv == "" {
		return fmt.Errorf("CodeValue is empty")
	}
	for _, pr := range c.storage {
		if pr.CodeValue == cv {
			return fmt.Errorf("CodeValue already exists")
		}
	}
	return nil
}

func (c *ControllerProducts) ValidateExpiration(expiration string) error {
	if expiration == "" {
		return fmt.Errorf("Expiration is empty")
	}
	date, err := time.Parse("02/01/2006", expiration)
	fmt.Println("Date parsed: ", date)
	if err != nil {
		return fmt.Errorf("Expiration is an invalid date")
	}
	return nil
}

func writeErrorContent(w http.ResponseWriter, msg string, status int) {
	body := &ResponseBodyProduct{
		Message: msg,
		Data:    nil,
		Error:   true,
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (c *ControllerProducts) CreateProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody RequestBodyProduct

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			writeErrorContent(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if reqBody.Name == "" || reqBody.Quantity == 0 || reqBody.Price == 0 ||
			reqBody.CodeValue == "" || reqBody.Expiration == "" {
			writeErrorContent(w, "Invalid values", http.StatusUnprocessableEntity)
			return
		}

		err := c.ValidateCodeValue(reqBody.CodeValue)
		if err != nil {
			writeErrorContent(w, err.Error(), http.StatusConflict)
			return
		}

		err = c.ValidateExpiration(reqBody.Expiration)
		if err != nil {
			writeErrorContent(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		prod := &Product{
			ID:          len(c.storage) + 1,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		c.storage[prod.ID] = prod

		respBody := &ResponseBodyProduct{
			Message: "Product created",
			Data:    prod,
			Error:   false,
		}

		w.WriteHeader(http.StatusCreated) // status 201
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respBody)
	}
}

func (c *ControllerProducts) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {

		}

		id, err := strconv.Atoi(idParam)
		if err != nil || id < 0 {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		product := c.storage[id]

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
		return
	}
}
