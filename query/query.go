package query

import (
	"github.com/graphql-go/graphql"
	"go-graphql/model"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
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

					id := p.Args["id"].(string)

					userId, e := objectid.FromHex(id)

					if e != nil {
						return nil, e
					}

					result, err := model.User{Id: userId}.Load()

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
