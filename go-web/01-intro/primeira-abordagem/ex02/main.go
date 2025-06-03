/*
	Exercício 2 - Manipulação do Body

Vamos criar um endpoint chamado /greetings. Com uma pequena estrutura com nome e sobrenome que, quando colada, deve
responder no texto "Hello + nome + sobrenome".

O ponto de extremidade deve ser um método POST
O package JSON deve ser usado para resolver o exercício.
A resposta deve seguir a seguinte estrutura: "Hello Andrea Rivas".
A estrutura deve ter a seguinte aparência:
{
	"firstName": "Andrea",
	"lastName": "Rivas".
}
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GreetingsRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	router := chi.NewRouter()

	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {

		var req GreetingsRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s %s!", req.FirstName, req.LastName)))
	})

	http.ListenAndServe(":8080", router)
}
