package query

import (
	"github.com/graphql-go/graphql"
	"go-graphql/model"
	"errors"
	"fmt"
)

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        model.UserType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, errors.New("invalid id")
					}

					user := &model.User{
						Id: int64(id),
					}

					result, err := user.Load()

					if err != nil {
						return nil, err
					}
					result.Password = ""

					return result, err
				},
			},

			"users": &graphql.Field{
				Type:        graphql.NewList(model.UserType),
				Description: "Get user list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 50,
					},
					"skip": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					auth := params.Context.Value("auth")

					fmt.Println("Auth", auth)

					limit := params.Args["limit"].(int)
					skip := params.Args["skip"].(int)

					users, err := model.Users(limit, skip)

					if err != nil {
						return nil, err
					}
					return users, err
				},
			},
			"countUsers": &graphql.Field{
				Type:        graphql.Int,
				Description: "Get user list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					count, err := model.CountUsers()

					if err != nil {
						return nil, err
					}
					return count, err
				},
			},
		},
	})
