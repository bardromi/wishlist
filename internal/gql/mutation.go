package gql

import "github.com/graphql-go/graphql"

func buildMutation(resolver Resolver, typeResolver TypeResolver) *graphql.Object {
	var mutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				//////////////////////////////////////////////////////////////////
				///////////////////////        USER        ///////////////////////
				//////////////////////////////////////////////////////////////////
				"SignUp": &graphql.Field{
					Type:        userType(typeResolver),
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
					Resolve: resolver.signUp,
				},
				"SignIn": &graphql.Field{
					Type:        userType(typeResolver),
					Description: "Sign in user",
					Args: graphql.FieldConfigArgument{
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.signIn,
				},
				"updateUser": &graphql.Field{
					Type:        userType(typeResolver),
					Description: "Update user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
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
					Resolve: resolver.userUpdateUser,
				},
				"deleteUser": &graphql.Field{
					Type:        userType(typeResolver),
					Description: "Delete user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.userDeleteUser,
				},

				//////////////////////////////////////////////////////////////////
				///////////////////////        WISH        ///////////////////////
				//////////////////////////////////////////////////////////////////
				"createWish": &graphql.Field{
					Type:        wishType(typeResolver),
					Description: "Create new wish",
					Args: graphql.FieldConfigArgument{
						"owner": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"title": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"price": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
					},
					Resolve: resolver.wishCreateWish,
				},
				"deleteWish": &graphql.Field{
					Type:        wishType(typeResolver),
					Description: "Delete wish by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: resolver.wishDeleteWish,
				},
				"updateWish": &graphql.Field{
					Type:        wishType(typeResolver),
					Description: "Update wish by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"title": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"price": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
					},
					Resolve: resolver.wishUpdateWish,
				},
			},
		},
	)

	return mutationType
}
