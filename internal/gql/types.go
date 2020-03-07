package gql

import (
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func wishType(typeResolver TypeResolver) *graphql.Object {
	var wishType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Wish",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"owner": &graphql.Field{
					Type:    userType,
					Resolve: typeResolver.getUserbyID,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Float,
				},
			},
		},
	)
	return wishType
}
