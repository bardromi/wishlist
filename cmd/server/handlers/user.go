package handlers

import (
	"fmt"
	"github.com/bardromi/wishlist/internal/platform/db"
	"github.com/bardromi/wishlist/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	MasterDB *db.DB
}

func (u *User) GetUser(c *gin.Context) {
	params := c.Request.URL.Query()
	usr, err := user.FindByEmail(u.MasterDB, params.Get("email"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"error": "no data for you",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": usr.Email,
	})
}

func (u *User) SignUp(c *gin.Context) {
	var err error
	var nu user.NewUser
	err = c.BindJSON(&nu)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	usr, err := user.SignUp(u.MasterDB, &nu)
	c.JSON(200, gin.H{
		"user": usr,
	})
}
