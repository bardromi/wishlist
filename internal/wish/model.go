package wish

import "time"

type Wish struct {
	ID        string
	OwnerID   string
	Title     string
	Price     string
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
