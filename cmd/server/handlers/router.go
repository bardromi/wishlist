package handlers

import (
	"github.com/bardromi/wishlist/internal/platform/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func API(masterDB *db.DB) http.Handler {
	router := gin.Default()

	u := User{
		MasterDB: masterDB,
	}

	router.GET("/", u.GetUser)
	router.POST("/", u.SignUp)

	return router
}
