package application

import (
	"estrutura-api/config"
	"estrutura-api/internal/application/routes"
	"net/http"
)

type Application struct {
	Router http.Handler
}

func NewApplication() *Application {
	config.Init()

	app := &Application{}
	app.Router = routes.NewRouter().MapRoutes()

	return app
}
