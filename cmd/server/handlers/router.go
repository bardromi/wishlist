package handlers

import (
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

	router.HandleFunc("/users/{id}", u.GetUser).Methods("GET")
	router.HandleFunc("/users", u.List).Methods("GET")
	router.HandleFunc("/signup", u.SignUp).Methods("POST")
	//router.HandleFunc("/signin", u.SignIn).Methods("POST")

	return router
}
