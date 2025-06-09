package main

import (
	"estrutura-api/internal/application"
	"net/http"
)

func main() {

	app := application.NewApplication()

	if err := http.ListenAndServe(":8080", app.Router); err != nil {
		panic(err)
	}
}
