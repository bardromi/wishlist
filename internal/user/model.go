package user

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
