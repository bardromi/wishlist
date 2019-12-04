package handlers

import (
	"github.com/bardromi/wishlist/internal/gql"
	"github.com/bardromi/wishlist/internal/mid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func API(db *sqlx.DB) http.Handler {
	//router := gin.Default()
	router := mux.NewRouter()

	gqlRoot := gql.NewRoot(db)

	graphql := GraphQL{
		GqlSchema: gqlRoot,
	}

	u := User{
		db: db,
	}

	router.Use(mid.Logging)

	router.HandleFunc("/graphql", graphql.GraphQL)
	router.HandleFunc("/users/{id}", u.GetUser).Methods("GET")
	router.Handle("/users", mid.Chain(http.HandlerFunc(u.List), mid.Authenticated)).Methods("GET")
	router.HandleFunc("/signup", u.SignUp).Methods("POST")
	router.HandleFunc("/signin", u.SignIn).Methods("POST")
	router.Handle("/signout", mid.Chain(http.HandlerFunc(u.SignOut), mid.Authenticated)).Methods("POST")

	return router
}
