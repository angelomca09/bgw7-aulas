package routes

import (
	"estrutura-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func ConfigProductRoutes(hd handler.ProductHandler, r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Post("/", hd.Create())
		r.Put("/{id}", hd.UpdateOrCreate())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
		r.Get("/consumer_price", hd.GetTotalPriceByList())
	})
}
