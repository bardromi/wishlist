package mid

import (
	"log"
	"net/http"

	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/bardromi/wishlist/internal/platform/web"
)

type MiddleWare func(http.Handler) http.Handler

func Chain(h http.Handler, middleware ...MiddleWare) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		cookie, err := r.Cookie("WishList")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := cookie.Value

		_, err = auth.ParseClaims(tknStr)

		if err != nil {
			web.RespondError(w, "User not Logged in", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}
