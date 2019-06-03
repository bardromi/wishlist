package user

import (
	"github.com/bardromi/wishlist/internal/platform/db"
)

func FindByEmail(dbConn *db.DB, email string) (user User, err error) {
	row := dbConn.FindByEmail(email)
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
