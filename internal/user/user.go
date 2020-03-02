package user

import (
	"database/sql"
	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"

	"errors"
	"github.com/jmoiron/sqlx"
)

var (
	// ErrNotFound abstracts the postgres not found error.
	ErrNotFound = errors.New("entity not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrValidateConfirmPassword occurs when password doesnt match
	ErrValidateConfirmPassword = errors.New("password confirmation failed")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")
)

func GetUserById(db *sqlx.DB, id string) (*User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrInvalidID
	}

	var u User
	const q = `SELECT * FROM users WHERE id = $1`

	if err := db.Get(&u, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &u, nil
}

func List(db *sqlx.DB) ([]User, error) {
	var users []User
	const q = `SELECT * FROM users`

	if err := db.Select(&users, q); err != nil {
		return nil, err
	}

	return users, nil
}

func SignUp(db *sqlx.DB, nu *NewUser) (*User, error) {
	if nu.Password != nu.PasswordConfirm {
		return nil, ErrValidateConfirmPassword
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nu.Password), 8)
	if err != nil {
		return nil, err
	}

	u := User{
		ID:           uuid.New().String(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}

	const q = "insert into users (ID, name, email, password, created_at) values ($1, $2, $3, $4, $5)"

	_, err = db.Exec(q, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func SignIn(db *sqlx.DB, now time.Time, email, password string) (auth.Claims, error) {
	const q = `SELECT * FROM users WHERE email = $1`
	var u User

	if err := db.Get(&u, q, email); err != nil {
		// Normally we would return ErrNotFound in this scenario but we do not want
		// to leak to an unauthenticated user which emails are in the system.
		if err == sql.ErrNoRows {
			return auth.Claims{}, ErrAuthenticationFailure
		}

		return auth.Claims{}, err
	}

	// Compare the provided password with the saved hash. Use the bcrypt
	// comparison function so it is cryptographically secure.
	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)); err != nil {
		return auth.Claims{}, ErrAuthenticationFailure
	}

	// If we are this far the request is valid. Create some claims for the user
	// and generate their token.
	claims := auth.NewClaims(u.Email, now, time.Hour)

	return claims, nil
}
