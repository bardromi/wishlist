package user

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
