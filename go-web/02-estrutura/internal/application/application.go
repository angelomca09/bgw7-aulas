package application

import (
	"estrutura-api/internal/application/routes"
	"estrutura-api/internal/handler"
	"estrutura-api/internal/repository"
	"estrutura-api/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Router         http.Handler
	productHandler *handler.ProductHandler
}

func NewApplication() *Application {
	app := &Application{}
	rt := chi.NewRouter()

	app.CreateHandlers()
	app.CreateRoutes(rt)

	return app
}

func (a *Application) CreateHandlers() {
	prodRepo := repository.NewProductJSONRepository()
	prodService := service.NewProductService(prodRepo)
	prodHandler := handler.NewProductHandler(prodService)

	a.productHandler = prodHandler
}

func (a *Application) CreateRoutes(r chi.Router) {
	routes.ConfigProductRoutes(*a.productHandler, r)

	a.Router = r
}
