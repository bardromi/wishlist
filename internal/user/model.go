package user

import "time"

// User represents user in database
type User struct {
	ID           string    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash []byte    `db:"password" json:"-"`
	DateCreated  time.Time `db:"date_created" json:"date_created"`
	DateUpdated  time.Time `db:"date_updated" json:"date_updated"`
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm" validate:"eqfield=Password"`
}

// UpdateUser defines what information may be provided to modify an existing
// User. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types but we make exceptions around
// marshalling/unmarshalling.
type UpdateUser struct {
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"password_confirm" validate:"omitempty,eqfield=Password"`
}
