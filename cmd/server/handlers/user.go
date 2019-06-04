package handlers

import (
	"fmt"
	"github.com/bardromi/wishlist/internal/platform/db"
	"github.com/bardromi/wishlist/internal/user"
	"github.com/gin-gonic/gin"
	"log"
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

func (u *User) CreateUser(c *gin.Context) {
	var err error
	var nu user.NewUser
	err = c.BindJSON(&nu)
	if err != nil {
		fmt.Println(err)
	}

	usr, err := user.CreateUser(u.MasterDB, &nu)
	fmt.Println(usr)
}
