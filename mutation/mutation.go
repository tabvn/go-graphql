package mutation

import (
	"github.com/graphql-go/graphql"
	"go-graphql/model"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{

		"createUser": &graphql.Field{
			Type:        model.UserType,
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"first_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"last_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				user := model.User{
					FirstName: params.Args["first_name"].(string),
					LastName:  params.Args["last_name"].(string),
					Email:     params.Args["email"].(string),
					Password:  params.Args["password"].(string),
				}

				result, err := user.Create()

				if err != nil {
					return nil, err
				}

				result.Password = ""

				return result, err

			},
		},

		"updateUser": &graphql.Field{
			Type:        model.UserType,
			Description: "Update user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"first_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"last_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				//id := params.Args["id"]

				user := model.User{
					Id:        1,
					FirstName: params.Args["first_name"].(string),
					LastName:  params.Args["last_name"].(string),
					Email:     params.Args["email"].(string),
					Password:  params.Args["password"].(string),

				}

				result, err := user.Update()

				return result, err

			},
		},
	},
})
