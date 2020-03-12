package user

import (
	"database/sql"
	"time"

	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

var (
	// ErrNotFound abstracts the postgres not found error.
	ErrNotFound = errors.New("entity not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrValidateConfirmPassword occurs when password doesn't match
	ErrValidateConfirmPassword = errors.New("password confirmation failed")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Retrieve gets the specified user from the database.
func Retrieve(db *sqlx.DB, id string) (*User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrInvalidID
	}

	var u User
	const q = `
	SELECT *
	FROM users
	WHERE id = $1`

	if err := db.Get(&u, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, errors.Wrapf(err, "selecting user %q", id)
	}

	return &u, nil
}

// List retrieves a list of existing users from the database.
func List(db *sqlx.DB) ([]User, error) {
	var users []User
	const q = `SELECT * FROM users`

	if err := db.Select(&users, q); err != nil {
		return nil, errors.Wrap(err, "selecting users")
	}

	return users, nil
}

// Create inserts a new user into the database.
// if want to use with graphql and reserve the package oriented design should get fields of new user and not new user
func Create(db *sqlx.DB, nu *NewUser) (*User, error) {
	if nu.Password != nu.PasswordConfirm {
		return nil, ErrValidateConfirmPassword
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nu.Password), 8)
	if err != nil {
		return nil, errors.Wrap(err, "generating password hash")
	}

	u := User{
		ID:           uuid.New().String(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hashedPassword,
		DateCreated:  time.Now(),
		DateUpdated:  time.Now(),
	}

	const q = `
	INSERT INTO users 
	(ID, name, email, password, date_created, date_updated)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = db.Exec(q, u.ID, u.Name, u.Email, u.PasswordHash, u.DateCreated, u.DateUpdated)
	if err != nil {
		return nil, errors.Wrap(err, "inserting user")
	}

	return &u, nil
}

// Update replaces a user document in the database.
func Update(db *sqlx.DB, id string, upd UpdateUser) error {
	u, err := Retrieve(db, id)
	if err != nil {
		return err
	}

	if upd.Name != nil {
		u.Name = *upd.Name
	}
	if upd.Email != nil {
		u.Email = *upd.Email
	}
	if upd.Password != nil {
		// Salt and hash the password using the bcrypt algorithm
		// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*upd.Password), 8)
		if err != nil {
			return errors.Wrap(err, "generating password hash")
		}
		u.PasswordHash = hashedPassword
	}

	u.DateUpdated = time.Now()

	const q = `
	UPDATE users SET
	"name" = $2
	"email" = $3
	"password = $4
	"date_updated" = $5
	WHERE id = $1`

	_, err = db.Exec(q, id, u.Name, u.Email, u.PasswordHash, u.DateUpdated)
	if err != nil {
		return errors.Wrap(err, "updating user")
	}

	return nil
}

// Authenticate finds a user by their email and verifies their password. On
// success it returns a Claims value representing this user. The claims can be
// used to generate a token for future authentication.
func Authenticate(db *sqlx.DB, now time.Time, email, password string) (auth.Claims, error) {
	const q = `
	SELECT * 
	FROM users
	WHERE email = $1`

	var u User

	if err := db.Get(&u, q, email); err != nil {
		// Normally we would return ErrNotFound in this scenario but we do not want
		// to leak to an unauthenticated user which emails are in the system.
		if err == sql.ErrNoRows {
			return auth.Claims{}, ErrAuthenticationFailure
		}

		return auth.Claims{}, errors.Wrapf(err, "selecting single user %s", email)
	}

	// Compare the provided password with the saved hash. Use the bcrypt
	// comparison function so it is cryptographically secure.
	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)); err != nil {
		return auth.Claims{}, ErrAuthenticationFailure
	}

	// If we are this far the request is valid. Create some claims for the user
	// and generate their token.
	claims := auth.NewClaims(u.ID, u.Email, now, time.Hour)

	return claims, nil
}

// Delete removes a user from the database.
func Delete(db *sqlx.DB, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidID
	}

	const q = `DELETE FROM users WHERE id = $1`

	if _, err := db.Exec(q, id); err != nil {
		return errors.Wrapf(err, "deleting user %s", id)
	}

	return nil
}
