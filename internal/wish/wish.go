package wish

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Create adds a Wish to the database. It returns the created Wish.
func Create(db *sqlx.DB, nw *NewWish) (*Wish, error) {

	wish := Wish{
		OwnerID:   nw.OwnerID,
		Title:     nw.Title,
		Price:     nw.Price,
		CreatedAt: time.Now(),
	}

	const q = `
	INSERT INTO wishes
	(owner_id, title, price, created_at)
	VALUES($1, $2, $3, $4)
	`

	row, err := db.Exec(q, wish.OwnerID, wish.Title, wish.Price, wish.CreatedAt)
	id, err := row.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "inserting wish")
	}

	wish.ID = id

	return &wish, nil
}
