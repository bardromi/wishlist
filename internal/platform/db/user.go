package db

import (
	"database/sql"
	"time"
)

func (db *DB) FindByEmail(email string) *sql.Row {
	return db.database.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE email = $1", email)
}

func (db *DB) CreateUser(name string, email string, password string, confirmPassword string) (*sql.Row, error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "insert into users (name, email, password, created_at) values ($1, $2, $3, $4) returning id, name, email, password, created_at"
	stmt, err := db.database.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	queryRow := stmt.QueryRow(name, email, password, time.Now())
	return queryRow, err
}
