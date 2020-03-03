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
func (r *Resolver) userGetUserByID(p graphql.ResolveParams) (interface{}, error) {
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
func (r *Resolver) userList(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	users, err := user.List(r.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// SignUp graphql connector to create user
func (r *Resolver) signUp(p graphql.ResolveParams) (interface{}, error) {
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
func (r *Resolver) signIn(p graphql.ResolveParams) (interface{}, error) {
	return nil, errors.New("not Implemented")
}

func (r *Resolver) wishCreateWish(p graphql.ResolveParams) (interface{}, error) {

	// owner,err := r.userGetUserByID(p.Args)

	// nw := wish.NewWish{
	// 	Title: p.Args["title"].(string),
	// 	Price: p.Args["price"].(float64),
	// }
	return nil, errors.New("not Implemented")
}
