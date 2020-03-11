package gql

import (
	"context"
	"errors"
	"fmt"

	"github.com/bardromi/wishlist/internal/platform/auth"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

type key string

// NewRoot create scheme root includes Query and Mutation
func NewRoot(db *sqlx.DB) *graphql.Schema {
	resolver := Resolver{db: db}
	typeResolver := TypeResolver{db: db}

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    buildQuery(resolver, typeResolver),
			Mutation: buildMutation(resolver, typeResolver),
		},
	)

	return &schema
}

// ExecuteQuery runs our graphql queries
func ExecuteQuery(query string, schema graphql.Schema, claims auth.Claims) (*graphql.Result, error) {
	const token key = "token"
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.WithValue(context.Background(), token, claims),
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
		//Todo: return all errors
		return nil, errors.New(result.Errors[0].Message)
	}
	return result, nil
}
