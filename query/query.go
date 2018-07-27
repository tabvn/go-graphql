package query

import (
	"github.com/graphql-go/graphql"
	"go-graphql/model"
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
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					id := p.Args["id"].(int64)

					result, err := model.User{Id: id}.Load()

					return result, err
				},
			},

			"users": &graphql.Field{
				Type:        graphql.NewList(model.UserType),
				Description: "Get user list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	})
