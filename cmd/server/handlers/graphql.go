package handlers

import (
	"encoding/json"
	"github.com/bardromi/wishlist/internal/gql"
	"github.com/bardromi/wishlist/internal/platform/web"
	"github.com/graphql-go/graphql"
	"net/http"
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
		http.Error(w, "Must provide graphql query in request body", 400)
		return
	}

	var rBody reqBody
	// Decode the request body into rBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		web.RespondError(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	// Execute graphql query
	result, err := gql.ExecuteQuery(rBody.Query, *g.GqlSchema)

	if err != nil {
		web.RespondError(w, "graphql could not process the request", http.StatusInternalServerError)
		return
	}

	web.Respond(w, result, http.StatusOK)
}
