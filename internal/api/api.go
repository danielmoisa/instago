package api

import (
	"log"
	"net/http"
	"time"

	"github.com.danielmoisa/instago/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config Config
	Store  store.Storage
}

type Config struct {
	Addr string
	Db   DbConfig
}

type DbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *Application) Run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 50,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running at %s", app.Config.Addr)

	return srv.ListenAndServe()
}
