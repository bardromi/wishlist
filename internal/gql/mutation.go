package gql

import "github.com/graphql-go/graphql"

func buildMutation(resolver Resolver) *graphql.Object {
	var mutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"SignUp": &graphql.Field{
					Type:        userType,
					Description: "Sign up new user",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"passwordConfirm": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.SignUp,
				},
				"SignIn": &graphql.Field{
					Type:        userType,
					Description: "Sign in user",
					Args: graphql.FieldConfigArgument{
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.SignIn,
				},
				// Todo: Implement
				"updateUser": &graphql.Field{
					Type:        userType,
					Description: "Update user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
				},
				// Todo: Implemet
				"deleteUser": &graphql.Field{
					Type:        userType,
					Description: "Delete user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
				},
			},
		},
	)

	return mutationType
}
