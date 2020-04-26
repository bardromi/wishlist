package gql

import "github.com/graphql-go/graphql"

func buildQuery(resolver Resolver, typeResolver TypeResolver) *graphql.Object {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
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
				"users": &graphql.Field{
					Type:        graphql.NewList(userType(typeResolver)),
					Description: "Get list of uers",
					Resolve:     resolver.userList,
				},
				"wish": &graphql.Field{
					Type:        wishType(typeResolver),
					Description: "Get wish by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: resolver.wishGetWishByID,
				},
				"wishes": &graphql.Field{
					Type:        graphql.NewList(wishType(typeResolver)),
					Description: "Get list of wishes",
					Resolve:     resolver.wishList,
				},
			},
		})

	return queryType
}
