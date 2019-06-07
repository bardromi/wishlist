package user

import (
	"database/sql"
	"errors"
	"github.com/bardromi/wishlist/internal/platform/db"
)

var (
	// ErrNotFound abstracts the postgres not found error.
	ErrNotFound = errors.New("entity not found")
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

func CreateUser(dbConn *db.DB, nu *NewUser) (*User, error) {
	var err error
	var user = User{}
	row, err := dbConn.CreateUser(nu.Name, nu.Email, nu.Password, nu.PasswordConfirm)
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
