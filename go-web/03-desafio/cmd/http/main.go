package main

import (
	"app/cmd/http/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "docs/db/tickets.csv",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db, err := loader.NewLoaderTicketCSV(a.dbFile).Load()
	if err != nil {
		return
	}
	rp := repository.NewRepositoryTicketMap(db, len(db))
	// service ...
	sv := service.NewServiceTicketDefault(rp)
	// handler ...
	hd := handler.NewTicketsHandlerDefault(sv)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	(*a).rt.Route("/ticket", func(r chi.Router) {
		r.Get("/", hd.GetAllTickets())
		r.Get("/getByCountry/{dest}", hd.GetTicketsByCountry())
		r.Get("/getAverage/{dest}", hd.GetAverageTicketsByCountry())
	})

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
