package user

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
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
