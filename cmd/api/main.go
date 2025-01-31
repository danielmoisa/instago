package main

import (
	"log"

	"github.com.danielmoisa/instago/internal/api"
)

func main() {
	cfg := &api.Config{
		Addr: ":8080",
	}
	app := &api.Application{
		Config: *cfg,
	}

	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
