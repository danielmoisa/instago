package main

import (
	"log"

	"github.com.danielmoisa/instago/internal/api"
	"github.com.danielmoisa/instago/internal/db"
	"github.com.danielmoisa/instago/internal/env"
	"github.com.danielmoisa/instago/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	cfg := api.Config{
		Addr: ":" + env.GetString("PORT", "3000"),
		Db: api.DbConfig{
			Addr:         env.GetString("DB_CONN", ""),
			MaxOpenConns: env.GetInteger("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInteger("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.Db.Addr, cfg.Db.MaxOpenConns, cfg.Db.MaxIdleConns, cfg.Db.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection established")

	repository := store.NewRepository(db)

	app := &api.Application{
		Config: cfg,
		Store:  repository,
	}

	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
