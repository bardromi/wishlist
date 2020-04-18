package gql

import "github.com/graphql-go/graphql"

func buildQuery(resolver Resolver, typeResolver TypeResolver) *graphql.Object {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				/* Get (read) single product by id
				   http://localhost:8080/user?query={user(id:1){id,name,email}}
				*/
				"user": &graphql.Field{
					Type:        userType(typeResolver),
					Description: "Get user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.userGetUserByID,
				},
				/* Get (read) users list
				   http://localhost:8080/product?query={list{id,name,info,price}}
				*/
				"users": &graphql.Field{
					Type:        graphql.NewList(userType(typeResolver)),
					Description: "Get list of uers",
					Resolve:     resolver.userList,
				},
				// Get (read) single wish by id
				"wish": &graphql.Field{
					Type:        wishType(typeResolver),
					Description: "Get wish by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
				},
				// Get (read) wishes list
				"wishes": &graphql.Field{
					Type:        graphql.NewList(wishType(typeResolver)),
					Description: "Get list of wishes",
				},
			},
		})

	return queryType
}
