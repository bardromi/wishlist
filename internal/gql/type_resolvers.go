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

func (rt *TypeResolver) getUserByID(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id := p.Source.(*wish.Wish).OwnerID
	users, err := user.Retrieve(rt.db, id)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (rt *TypeResolver) getUserWishes(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id := p.Source.(*user.User).ID
	wishes, err := wish.GetWishesByUserID(rt.db, id)
	if err != nil {
		return nil, err
	}

	return wishes, nil
}
