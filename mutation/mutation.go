package mutation

import (
	"github.com/graphql-go/graphql"
	"go-graphql/model"
	"errors"
	"fmt"
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
					Type: graphql.NewNonNull(graphql.Int),
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

				id, ok := params.Args["id"].(int)
				if !ok {
					return nil, errors.New("invalid id")
				}
				user := model.User{
					Id:        int64(id),
					FirstName: params.Args["first_name"].(string),
					LastName:  params.Args["last_name"].(string),
					Email:     params.Args["email"].(string),
					Password:  params.Args["password"].(string),
				}

				result, err := user.Update()

				if err != nil {
					return nil, err
				}
				return result, err

			},
		},

		"login": &graphql.Field{
			Type:        model.LoginType,
			Description: "Login",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				email := params.Args["email"].(string)
				password := params.Args["password"].(string)

				token, user, err := model.LoginUser(email, password)

				fmt.Println("error", token, user, err)

				type UserResult struct {
					user *model.User
				}

				type LoginResult struct {
					*model.Token
					UserResult `json:"user"`
				}
				result := LoginResult{
					token,
					UserResult{
						user: user,
					},
				}

				fmt.Println("result", result.user)

				return result, nil

			},
		},
	},
})
