package gql

import (
	"errors"

	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/bardromi/wishlist/internal/user"
	"github.com/bardromi/wishlist/internal/wish"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *sqlx.DB
}

var (
	// ErrValidationFailed abstracts the postgres not found error.
	ErrValidationFailed = errors.New("one or more of parameters are invalid")
)

// UserGetUserByID graphql connector to get user by id
func (r *Resolver) userGetUserByID(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id, ok := p.Args["id"].(string)
	if ok {
		user, err := user.Retrieve(r.db, id)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, ErrValidationFailed
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
	name, okName := p.Args["name"].(string)
	email, okEmail := p.Args["email"].(string)
	password, okPassword := p.Args["password"].(string)
	passwordConfirm, okPasswordConfirm := p.Args["passwordConfirm"].(string)

	if okName && okEmail && okPassword && okPasswordConfirm {
		nu := user.NewUser{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
		}

		usr, err := user.Create(r.db, &nu)
		if err != nil {
			return nil, err
		}

		return usr, nil
	}

	return nil, ErrValidationFailed
}

// SignIn graphql connector to authenticate <<not implemented yet>>
func (r *Resolver) signIn(p graphql.ResolveParams) (interface{}, error) {
	return nil, errors.New("not Implemented")
}

func (r *Resolver) userDeleteUser(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id, ok := p.Args["id"].(string)
	if ok {

		userFromDB, err := user.Retrieve(r.db, id)
		if err != nil {
			return nil, err
		}

		err = user.Delete(r.db, id)
		if err != nil {
			return nil, err
		}

		return userFromDB, nil
	}

	return nil, ErrValidationFailed
}

func (r *Resolver) userUpdateUser(p graphql.ResolveParams) (interface{}, error) {
	id, okID := p.Args["id"].(string)
	if !okID {
		return nil, ErrValidationFailed
	}

	var pName, pEmail, pPassword, pPasswordConfirm *string
	// maybe unmarshal into userupdate
	name, okName := p.Args["name"].(string)
	if okName {
		pName = &name
	}

	email, okEmail := p.Args["email"].(string)
	if okEmail {
		pEmail = &email
	}

	password, okPassword := p.Args["password"].(string)
	if okPassword {
		pPassword = &password
	}

	passwordConfirm, okPassordConfirm := p.Args["passwordConfirm"].(string)
	if okPassordConfirm {
		pPasswordConfirm = &passwordConfirm
	}

	updateUser := user.UpdateUser{
		Name:            pName,
		Email:           pEmail,
		Password:        pPassword,
		PasswordConfirm: pPasswordConfirm,
	}

	user, err := user.Update(r.db, id, updateUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) wishGetWishByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)

	if !ok {
		return nil, ErrValidationFailed
	}

	wish, err := wish.GetWishesByUserID(r.db, id)
	if err != nil {
		return nil, err
	}

	return wish, nil
}

func (r *Resolver) wishList(p graphql.ResolveParams) (interface{}, error) {
	wishes, err := wish.List(r.db)

	if err != nil {
		return nil, err
	}

	return wishes, nil
}

func (r *Resolver) wishCreateWish(p graphql.ResolveParams) (interface{}, error) {

	claim := p.Context.Value("token").(auth.Claims)

	if claim.UserID == "" {
		return nil, errors.New("Cannot add wish without owner")
	}

	// Todo: type assertion
	nw := wish.NewWish{
		OwnerID: claim.UserID,
		Title:   p.Args["title"].(string),
		Price:   p.Args["price"].(float64),
	}

	wish, err := wish.Create(r.db, &nw)
	if err != nil {
		return nil, err
	}

	return wish, nil
}
