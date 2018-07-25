package model

import (
	"time"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/graphql-go/graphql"
)

type User struct {
	Id        objectid.ObjectID `json:"id"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Created   time.Time         `json:"created"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"created": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
