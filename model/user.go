package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"context"
	"log"
	"go-graphql/db"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/bson"
)

type User struct {
	Id        string    `json:"_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Created   time.Time `json:"created"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: graphql.String,
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
func Create(user User) (User, error) {

	collection := db.Collection("users")

	id := objectid.New()

	user.Id = string(id.Hex())

	_, err := collection.InsertOne(context.Background(), bson.NewDocument(
		bson.EC.ObjectID("_id", id),
		bson.EC.String("first_name", user.FirstName),
		bson.EC.String("last_name", user.LastName),
		bson.EC.String("email", user.Email),
		bson.EC.String("password", user.Password),
		bson.EC.Time("created", user.Created),
	))
	if err != nil {
		log.Fatal(err)

		return user, err
	}

	return user, nil
}
