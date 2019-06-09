package handlers

import (
	"github.com/bardromi/wishlist/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type User struct {
	db *sqlx.DB
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetUser(c *gin.Context) {
	usr, err := user.GetUserById(u.db, c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"error": "no data for you",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": usr,
	})
}

func (u *User) List(c *gin.Context) {
	usrs, err := user.List(u.db)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"users": usrs,
	})

}

func (u *User) SignUp(c *gin.Context) {
	var nu user.NewUser

	err := c.BindJSON(&nu)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	usr, err := user.SignUp(u.db, &nu)
	c.JSON(200, gin.H{
		"user": usr,
	})
}

func (u *User) SignIn(c *gin.Context) {
	var login Login

	err := c.BindJSON(&login)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	usr, err := user.SignIn(u.db, login.Email, login.Password)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"user": usr,
	})
}
