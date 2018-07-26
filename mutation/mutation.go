package mutation

import (
	"github.com/graphql-go/graphql"
	"go-graphql/model"
	"time"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
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
					Created:   time.Now(),
					Updated:   time.Now(),
				}

				result, err := user.Create()

				return result, err

			},
		},

		"updateUser": &graphql.Field{
			Type:        model.UserType,
			Description: "Update user",
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
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
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				id, err := objectid.FromHex(params.Args["_id"].(string))

				user := model.User{
					Id:        id,
					FirstName: params.Args["first_name"].(string),
					LastName:  params.Args["last_name"].(string),
					Email:     params.Args["email"].(string),
					Password:  params.Args["password"].(string),
					Created:   time.Now(),
					Updated:   time.Now(),
				}

				result, err := user.Update()

				return result, err

			},
		},
	},
})
