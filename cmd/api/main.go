package main

import (
	"log"

	"github.com.danielmoisa/instago/internal/api"
	"github.com.danielmoisa/instago/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &api.Config{
		Addr: ":" + env.GetString("PORT", "3000"),
	}
	app := &api.Application{
		Config: *cfg,
	}

	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
