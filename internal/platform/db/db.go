package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	database *sql.DB
}

func New(kind string, connection string) (*DB, error) {
	ses, err := sql.Open(kind, connection)

	if err != nil {
		log.Fatal(err)
	}
	if err = ses.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("##### Postgres Database initialized #####")
	db := DB{
		database: ses,
	}

	return &db, err
}

// Close closes a DB value being used with MongoDB.
func (db *DB) Close() {
	db.Close()
}
