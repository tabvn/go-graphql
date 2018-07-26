package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"go-graphql/db"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"go-graphql/scalar"
	"github.com/mongodb/mongo-go-driver/bson"
)

type User struct {
	Id        objectid.ObjectID `json:"_id" bson:"_id"`
	Email     string            `json:"email" bson:"email"`
	Password  string            `json:"password" bson:"password"`
	FirstName string            `json:"first_name" bson:"first_name"`
	LastName  string            `json:"last_name" bson:"last_name"`
	Created   time.Time         `json:"created" bson:"created"`
	Updated   time.Time         `json:"updated" bson:"updated"`
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

	id := objectid.New()

	u.Id = id

	_, err := db.Create("users", u)

	return u, err
}

func (u User) Update() (User, error) {

	type Filter struct {
		Id objectid.ObjectID `bson:"_id"`
	}

	update := bson.NewDocument(
		bson.EC.SubDocumentFromElements("$set",
			bson.EC.String("email", u.Email),
			bson.EC.String("first_name", u.FirstName),
			bson.EC.String("last_name", u.LastName),
			bson.EC.String("password", u.Password),
			bson.EC.Time("updated", time.Now()),
		))

	_, err := db.Update("users", Filter{Id: u.Id}, update)

	return u, err
}
