package routes

import (
	"estrutura-api/internal/handler"
	"estrutura-api/internal/repository"
	"estrutura-api/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func buildProductRoutes() http.Handler {
	r := chi.NewRouter()

	db := repository.NewProductJSONRepository("docs/db/products.json")
	s := service.NewProductService(db)
	hd := handler.NewProductHandler(s)

	r.Get("/", hd.GetAll())
	r.Get("/{id}", hd.GetById())
	r.Post("/", hd.Create())
	r.Put("/{id}", hd.UpdateOrCreate())
	r.Patch("/{id}", hd.Update())
	r.Delete("/{id}", hd.Delete())
	r.Get("/consumer_price", hd.GetTotalPriceByList())

	return r
}
