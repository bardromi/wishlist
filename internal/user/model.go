package user

import "time"

// User represents user in database
type User struct {
	ID           string    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash []byte    `db:"password" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
