package main

import (
	"github.com/bardromi/wishlist/cmd/server/handlers"
	"github.com/bardromi/wishlist/internal/platform/config"
	"github.com/bardromi/wishlist/internal/platform/database"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	// =========================================================================
	// Start Database

	db, err := database.Open(database.Config{
		User:       cfg.DatabaseConfiguration.User,
		Password:   cfg.DatabaseConfiguration.Password,
		Host:       cfg.DatabaseConfiguration.Host,
		Name:       cfg.DatabaseConfiguration.Name,
		DisableTLS: cfg.DatabaseConfiguration.DisableTLS,
	})
	if err != nil {
		log.Panic("Error initialize db", err)
	}

	defer func() {
		log.Printf("main : Database Stopping : %s", cfg.DatabaseConfiguration.Host)
		log.Fatal(db.Close())
	}()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handlers.API(db),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server started, Listening on port", s.Addr)
	log.Fatal(s.ListenAndServe())
}
