package gql

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

func NewRoot(db *sqlx.DB) *graphql.Schema {
	resolver := Resolver{db: db}

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    buildQuery(resolver),
			Mutation: buildMutation(resolver),
		},
	)

	return &schema
}

func ExecuteQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
		//Todo: return all errors
		return nil, errors.New(result.Errors[0].Message)
	}
	return result, nil
}
