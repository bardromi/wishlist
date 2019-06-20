package handlers

import (
	"encoding/json"
	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/bardromi/wishlist/internal/platform/web"
	"github.com/bardromi/wishlist/internal/user"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
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
		web.RespondError(w, "User not found", http.StatusOK)
		return
	}

	web.Respond(w, usr, http.StatusOK)
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	usrs, err := user.List(u.db)
	if err != nil {
		web.RespondError(w, "There are no users", http.StatusOK)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	web.Respond(w, usrs, http.StatusOK)
}

func (u *User) SignUp(w http.ResponseWriter, r *http.Request) {
	var nu user.NewUser
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		web.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr, err := user.SignUp(u.db, &nu)

	web.Respond(w, usr, http.StatusOK)
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	var login Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		web.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, err := user.SignIn(u.db, time.Now(), login.Email, login.Password)
	if err != nil {
		web.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tkn, err := auth.GenerateToken(claims)
	if err != nil {
		web.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "WishList",
		Value:    tkn,
		HttpOnly: true,
		Expires:  time.Unix(claims.ExpiresAt, 0),
	}
	http.SetCookie(w, &cookie)

	web.Respond(w, map[string]string{"email": claims.Email}, http.StatusOK)
}

//// respondWithJSON write json response format
//func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
//	response, _ := json.Marshal(payload)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//
//	// Send the result back to the client.
//	if _, err := w.Write(response); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//// respondWithError return error message
//func respondWithError(w http.ResponseWriter, code int, msg string) {
//	respondWithJSON(w, code, map[string]string{"message": msg})
//}
