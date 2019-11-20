package auth

import (
	"errors"
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
func NewClaims(email string, now time.Time, expires time.Duration) Claims {
	c := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Subject:  "wishList",
			IssuedAt: now.Unix(),
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: now.Add(expires).Unix(),
		},
	}

	return c
}

func ParseClaims(tknStr string) (Claims, error) {
	// Initialize a new instance of `Claims`
	var claims Claims

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, err
		}
		return claims, err
	}
	if !tkn.Valid {
		return claims, errors.New("invalid token")
	}

	return claims, nil
}

//// Valid is called during the parsing of a token.
//func (c Claims) Valid() error {
//	//for _, r := range c.Roles {
//	//	switch r {
//	//	case RoleAdmin, RoleUser: // Role is valid.
//	//	default:
//	//		return fmt.Errorf("invalid role %q", r)
//	//	}
//	//}
//	if err := c.StandardClaims.Valid(); err != nil {
//		return err
//	}
//	return nil
//}

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
