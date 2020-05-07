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

// UpdateWish defines what information may be provided to modify an existing
// Wish. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types but we make exceptions around
// marshalling/unmarshalling.
type UpdateWish struct {
	Title *string  `db:"title" json:"title"`
	Price *float64 `db:"price" json:"price"`
}
