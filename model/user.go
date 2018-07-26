package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"context"
	"log"
	"go-graphql/db"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"go-graphql/scalar"
)

type User struct {
	Id        objectid.ObjectID `json:"_id" bson:"_id"`
	Email string            `json:"email"`
	Password  string            `json:"password"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Created   time.Time         `json:"created"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{

			"_id": &graphql.Field{
				Type: scalar.ObjectIdType,
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

func (u User) Json() (User, error) {

	return u, nil

}
func (u User) Create() (User, error) {

	collection := db.Collection("users")

	id := objectid.New()

	u.Id = id

	_, err := collection.InsertOne(context.Background(), u)

	if err != nil {
		log.Fatal(err)

		return u, err
	}

	return u, nil
}
