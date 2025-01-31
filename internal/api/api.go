package api

import (
	"log"
	"net/http"
	"time"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func (app *Application) Run() error {
	mux := http.NewServeMux()

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
