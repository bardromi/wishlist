package mid

import (
	"log"
	"net/http"
)

type MiddleWare func(http.HandlerFunc) http.HandlerFunc

func Chain(h http.HandlerFunc, middleware ...MiddleWare) http.HandlerFunc {
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

func Stam(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("stam...")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	}
}
