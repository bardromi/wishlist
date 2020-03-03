package gql

import (
	"errors"

	"github.com/bardromi/wishlist/internal/user"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *sqlx.DB
}

// UserGetUserByID graphql connector to get user by id
func (r *Resolver) UserGetUserByID(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id, ok := p.Args["id"].(string)
	if ok {
		users, err := user.GetUserByID(r.db, id)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	return nil, nil
}

// UserList graphql connector to get all users
func (r *Resolver) UserList(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	users, err := user.List(r.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// SignUp graphql connector to create user
func (r *Resolver) SignUp(p graphql.ResolveParams) (interface{}, error) {
	nu := user.NewUser{
		Name:            p.Args["name"].(string),
		Email:           p.Args["email"].(string),
		Password:        p.Args["password"].(string),
		PasswordConfirm: p.Args["passwordConfirm"].(string),
	}

	usr, err := user.SignUp(r.db, &nu)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// SignIn graphql connector to authenticate <<not implemented yet>>
func (r *Resolver) SignIn(p graphql.ResolveParams) (interface{}, error) {
	return nil, errors.New("not Implemented")
}
