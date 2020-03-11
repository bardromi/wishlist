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
					Description: "Get users list",
					Resolve:     resolver.userList,
				},
			},
		})

	return queryType
}
