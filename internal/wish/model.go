package wish

import "time"

// Wish represents wish in database
type Wish struct {
	ID        string    `db:"id" json:"id"`
	OwnerID   string    `db:"owner_id" json:"ownerid"`
	Title     string    `db:"title" json:"title"`
	Price     int       `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
