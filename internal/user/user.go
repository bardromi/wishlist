package user

import (
	"database/sql"
	"errors"
	"github.com/bardromi/wishlist/internal/platform/db"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound abstracts the postgres not found error.
	ErrNotFound = errors.New("entity not found")

	ErrValidateConfirmPassword = errors.New("password confirmation failed")
)

func FindByEmail(dbConn *db.DB, email string) (*User, error) {
	var user = User{}
	row := dbConn.FindByEmail(email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, err
}

func SignUp(dbConn *db.DB, nu *NewUser) (*User, error) {
	var err error
	var user = User{}

	if nu.Password != nu.PasswordConfirm {
		return nil, ErrValidateConfirmPassword
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nu.Password), 8)

	row, err := dbConn.CreateUser(nu.Name, nu.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	if err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, err
}
