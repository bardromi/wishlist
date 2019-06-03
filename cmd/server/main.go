package main

import (
	"fmt"
	"github.com/bardromi/wishlist/cmd/server/handlers"
	"github.com/bardromi/wishlist/internal/platform/config"
	"github.com/bardromi/wishlist/internal/platform/db"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := config.LoadConfig()

	connection := fmt.Sprintf(`%s://%s:%s@%s/%s?sslmode=%s`, conf.Kind, conf.Username, conf.Password, conf.DatabaseConfiguration.Address, conf.DBName, conf.SslMode)

	log.Println("main : Started : Initialize Postgres")
	masterDB, err := db.New(conf.Kind, connection)
	if err != nil {
		log.Fatalf("main : Register DB : %v", err)
	}
	defer masterDB.Close()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handlers.API(masterDB),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = s.ListenAndServe()

	if err != nil {
		log.Fatalln("Server failed", err)
	}
}
