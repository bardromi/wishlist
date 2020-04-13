package wish

import "time"

// Wish represents wish in database
type Wish struct {
	ID          int64     `db:"id" json:"id"`
	OwnerID     string    `db:"owner_id" json:"ownerid"`
	Title       string    `db:"title" json:"title"`
	Price       float64   `db:"price" json:"price"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	DateUpdated time.Time `db:"date_updated" json:"date_updated"`
}

// NewWish contains information needed to create a new Wish.
type NewWish struct {
	OwnerID string  `json:"ownerid"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
}
