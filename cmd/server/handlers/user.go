package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bardromi/wishlist/internal/user"
	"github.com/gorilla/mux"
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

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr, err := user.GetUserById(u.db, vars["id"])

	if err != nil {
		log.Println(err.Error())
		respondWithError(w, http.StatusOK, "no data for you")
		return
	}

	respondWithJSON(w, http.StatusOK, usr)
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	usrs, err := user.List(u.db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	respondWithJSON(w, http.StatusOK, usrs)
}

func (u *User) SignUp(w http.ResponseWriter, r *http.Request) {
	var nu user.NewUser
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	usr, err := user.SignUp(u.db, &nu)

	respondWithJSON(w, http.StatusOK, usr)
}

//func (u *User) SignIn(w http.ResponseWriter, r *http.Request,) {
//	var login Login
//
//	err := c.BindJSON(&login)
//	if err != nil {
//		// If there is something wrong with the request body, return a 400 status
//		log.Println(err)
//		c.AbortWithStatus(http.StatusBadRequest)
//		return
//	}
//
//	claims, err := user.SignIn(u.db, login.Email, login.Password)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatus(http.StatusBadRequest)
//		return
//	}
//
//	tkn, err := auth.GenerateToken(claims)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//	cookie, err := c.Cookie("WishList")
//	//if Cookie doesnt exist or expired for current session create a new one
//	if err != nil {
//		cookie = "NotSet"
//		expires := int(claims.ExpiresAt - time.Now().Unix())
//		c.SetCookie("WishList", tkn, expires, "/", "localhost", false, true)
//	}
//	fmt.Printf("Cookie value: %s \n", cookie)
//
//	c.JSON(200, gin.H{
//		"Message": "Success",
//	})
//}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Send the result back to the client.
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// respondWithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
