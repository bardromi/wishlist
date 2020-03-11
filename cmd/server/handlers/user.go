package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/bardromi/wishlist/internal/platform/web"
	"github.com/bardromi/wishlist/internal/user"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// User represents the User API method handler set.
type User struct {
	db *sqlx.DB
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUser returns the specified user from the system.
func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr, err := user.GetUserByID(u.db, vars["id"])

	if err != nil {
		web.RespondError(w, "User not found", http.StatusOK)
		return
	}

	web.Respond(w, usr, http.StatusOK)
}

// List returns all the existing users in the system.
func (u *User) List(w http.ResponseWriter, r *http.Request) {
	usrs, err := user.List(u.db)
	if err != nil {
		web.RespondError(w, "There are no users", http.StatusOK)
		return
	}

	web.Respond(w, usrs, http.StatusOK)
}

// SignUp inserts a new user into the system.
func (u *User) SignUp(w http.ResponseWriter, r *http.Request) {
	var nu user.NewUser
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		web.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr, err := user.Create(u.db, &nu)
	if err != nil {
		web.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	web.Respond(w, usr, http.StatusOK)
}

// SignIn authenticate a user. It expects a request using
// Basic Auth with a user's email and password. It responds with a JWT.
func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	var login Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		web.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, err := user.Authenticate(u.db, time.Now(), login.Email, login.Password)
	if err != nil {
		web.RespondError(w, err.Error(), http.StatusUnauthorized)
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

// SignOut disconnect user from system
func (u *User) SignOut(w http.ResponseWriter, r *http.Request) {
	expiredCookie := http.Cookie{
		Name:     "WishList",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, &expiredCookie)
	web.Respond(w, "bye bye", http.StatusOK)
}
