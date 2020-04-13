package wish

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Create adds a Wish to the database. It returns the created Wish.
func Create(db *sqlx.DB, nw *NewWish) (*Wish, error) {

	wish := Wish{
		OwnerID:     nw.OwnerID,
		Title:       nw.Title,
		Price:       nw.Price,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	const q = `
	INSERT INTO wishes
	(owner_id, title, price, date_created, date_updated)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var wishID int64
	// I would use db.Exec (like in create user) but wish id (automatic generated) needed so QueryRow
	// is the solution here
	err = stmt.QueryRow(wish.OwnerID, wish.Title, wish.Price, wish.DateCreated, wish.DateUpdated).Scan(&wishID)
	if err != nil {
		return nil, errors.Wrap(err, "inserting wish")
	}
	// row, err := db.Exec(q, wish.OwnerID, wish.Title, wish.Price, wish.CreatedAt)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "inserting wish")
	// }

	wish.ID = wishID

	return &wish, nil
}

// GetWishesByUserID gets all user wishes from the database.
func GetWishesByUserID(db *sqlx.DB, id string) ([]*Wish, error) {
	var wishes []Wish

	const q = `
	SELECT * 
	FROM wishes
	WHERE owner_id =$1`

	if err := db.Select(&wishes, q, id); err != nil {
		return nil, errors.Wrapf(err, "selecting wishes by user %s", id)
	}

	wishesPointer := []*Wish{}

	for _, wish := range wishes {
		wishesPointer = append(wishesPointer, &wish)
	}

	return wishesPointer, nil
}
