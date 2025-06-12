package routes

import (
	"estrutura-api/internal/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type router struct{}

func NewRouter() *router {
	return &router{}
}

func (rt router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.JsonMiddleware)
	r.Use(middlewares.AuthMiddleware)

	r.Route("/products", func(r chi.Router) {
		r.Mount("/", buildProductRoutes())
	})

	return r
}
