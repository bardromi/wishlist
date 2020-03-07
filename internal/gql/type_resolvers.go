package gql

import (
	"github.com/bardromi/wishlist/internal/user"
	"github.com/bardromi/wishlist/internal/wish"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

// TypeResolver struct holds a connection to our database for gql types
type TypeResolver struct {
	db *sqlx.DB
}

func (rt *TypeResolver) getUserbyID(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id := p.Source.(*wish.Wish).OwnerID
	users, err := user.GetUserByID(rt.db, id)
	if err != nil {
		return nil, err
	}
	return users, nil
}
