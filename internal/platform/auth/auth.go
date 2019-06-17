package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"username"`
	jwt.StandardClaims
}

// NewClaims constructs a Claims value for the identified user. The Claims
// expire within a specified duration of the provided time. Additional fields
// of the Claims can be set after calling NewClaims is desired.
func NewClaims(email string, expires time.Duration) Claims {
	c := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Subject:  "wishList",
			IssuedAt: time.Now().Unix(),
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(expires).Unix(),
		},
	}

	return c
}

// Valid is called during the parsing of a token.
func (c Claims) Valid() error {
	//for _, r := range c.Roles {
	//	switch r {
	//	case RoleAdmin, RoleUser: // Role is valid.
	//	default:
	//		return fmt.Errorf("invalid role %q", r)
	//	}
	//}
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}
	return nil
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func GenerateToken(claims Claims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, err := tkn.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return str, nil
}

//// HasRole returns true if the claims has at least one of the provided roles.
//func (c Claims) HasRole(roles ...string) bool {
//	for _, has := range c.Roles {
//		for _, want := range roles {
//			if has == want {
//				return true
//			}
//		}
//	}
//	return false
//}
