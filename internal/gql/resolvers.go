package gql

import (
	"github.com/bardromi/wishlist/internal/user"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *sqlx.DB
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) UserGetUserById(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id, ok := p.Args["id"].(string)
	if ok {
		users, err := user.GetUserById(r.db, id)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	return nil, nil
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) UserList(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	users, err := user.List(r.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}
