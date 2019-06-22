package handlers

import (
	"github.com/bardromi/wishlist/internal/mid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func API(db *sqlx.DB) http.Handler {
	//router := gin.Default()
	router := mux.NewRouter()

	u := User{
		db: db,
	}

	//router.Use(mid.Logger)

	router.HandleFunc("/users/{id}", u.GetUser).Methods("GET")
	router.HandleFunc("/users", mid.Chain(u.List, mid.Logger)).Methods("GET")
	router.HandleFunc("/signup", u.SignUp).Methods("POST")
	router.HandleFunc("/signin", u.SignIn).Methods("POST")

	return router
}
