package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func API(db *sqlx.DB) http.Handler {
	router := gin.Default()

	u := User{
		db: db,
	}

	router.GET("/users/:id", u.GetUser)
	router.GET("/users", u.List)
	router.POST("/signup", u.SignUp)
	router.POST("/signin", u.SignIn)

	return router
}
