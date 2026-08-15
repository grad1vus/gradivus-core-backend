package main

import (
	"log"
	"net/http"

	"gradivus-core-backend/internal/config"
	"gradivus-core-backend/internal/database"
	"gradivus-core-backend/internal/routes"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	router := routes.New(db, cfg)

	log.Printf("Gradivus Core (Identity) API running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
