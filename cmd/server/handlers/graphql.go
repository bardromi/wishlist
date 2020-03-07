package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bardromi/wishlist/internal/gql"
	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/bardromi/wishlist/internal/platform/web"
	"github.com/graphql-go/graphql"
)

type GraphQL struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (g *GraphQL) GraphQL(w http.ResponseWriter, r *http.Request) {
	// Check to ensure query was provided in the request body
	if r.Body == nil {
		web.RespondError(w, "Must provide graphql query in request body", http.StatusBadRequest)
		return
	}

	var rBody reqBody
	// Decode the request body into rBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		web.RespondError(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	var claims auth.Claims
	cookie, claimsErr := r.Cookie("WishList")
	if claimsErr == nil {
		claims, _ = auth.ParseClaims(cookie.Value)
	}

	// Execute graphql query
	result, err := gql.ExecuteQuery(rBody.Query, *g.GqlSchema, claims)

	if err != nil {
		web.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Respond(w, result, http.StatusOK)
}
