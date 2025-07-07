package application

import (
	"app/internal/database"
	"app/internal/handler"
	"app/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewApplicationDefault creates a new default application.
func NewApplicationDatabase(addr string) (a *ApplicationDatabase) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	a = &ApplicationDatabase{
		rt:   defaultRouter,
		addr: defaultAddr,
	}
	return
}

// ApplicationDatabase is the default application.
type ApplicationDatabase struct {
	// rt is the router.
	rt *chi.Mux
	// addr is the address to listen.
	addr string
}

// TearDown tears down the application.
func (a *ApplicationDatabase) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDatabase) SetUp() (err error) {
	// dependencies
	// - db
	db := database.GetDB()
	// - repository
	rp := repository.NewRepositoryProductMysql(db)
	// - handler
	hd := handler.NewHandlerProduct(rp)

	wrp := repository.NewRepositoryWarehouseMysql(db)
	whd := handler.NewHandlerWarehouse(wrp)

	// router
	// - middlewares
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
		// GET warehouse/reportProducts
		r.Get("/warehouse/reportProducts", hd.GetTotalQuantityByWarehouse())
	})

	a.rt.Route("/warehouses", func(r chi.Router) {
		// GET /warehouses/{id}
		r.Get("/{id}", whd.GetById())
		// POST /warehouses
		r.Post("/", whd.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDatabase) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
