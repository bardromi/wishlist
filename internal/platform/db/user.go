package db

import (
	"database/sql"
)

func (db *DB) FindByEmail(email string) *sql.Row {
	return db.database.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE email = $1", email)
}
