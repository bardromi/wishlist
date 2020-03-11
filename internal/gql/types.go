package gql

import (
	"github.com/graphql-go/graphql"
)

func userType(typeResolver TypeResolver) *graphql.Object {
	var userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.FieldsThunk(func() graphql.Fields {
				return graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.String,
					},
					"name": &graphql.Field{
						Type: graphql.String,
					},
					"email": &graphql.Field{
						Type: graphql.String,
					},
					"wishes": &graphql.Field{
						Type:    graphql.NewList(wishType(typeResolver)),
						Resolve: typeResolver.getUserWishes,
					},
				}
			}),
		},
	)

	return userType
}

func wishType(typeResolver TypeResolver) *graphql.Object {
	var wishType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Wish",
			Fields: graphql.FieldsThunk(func() graphql.Fields {
				return graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.String,
					},
					"owner": &graphql.Field{
						Type:    userType(typeResolver),
						Resolve: typeResolver.getUserByID,
					},
					"title": &graphql.Field{
						Type: graphql.String,
					},
					"price": &graphql.Field{
						Type: graphql.Float,
					},
				}
			}),
		},
	)
	return wishType
}
