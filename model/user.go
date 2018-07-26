package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"go-graphql/db"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"go-graphql/scalar"
	"github.com/mongodb/mongo-go-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"context"
	"strings"
	"go-graphql/helper"
	"errors"
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
			"updated": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u User) Create() (User, error) {

	u, validateError := u.validateCreate()

	if validateError != nil {

		return u, validateError
	}

	id := objectid.New()

	u.Id = id

	// generate password
	password, e := HashPassword(u.Password)
	u.Password = password

	if e != nil {
		return u, e
	}

	_, err := db.Create("users", u)

	return u, err
}

func (u User) Update() (User, error) {

	type Filter struct {
		Id objectid.ObjectID `bson:"_id"`
	}

	if u.Password != "" {

	}

	update := bson.NewDocument(
		bson.EC.SubDocumentFromElements("$set",
			bson.EC.String("email", u.Email),
			bson.EC.String("first_name", u.FirstName),
			bson.EC.String("last_name", u.LastName),
			bson.EC.Time("updated", time.Now()),
		))

	_, err := db.Update("users", Filter{Id: u.Id}, update)

	return u, err
}

func (u User) Load() (User, error) {

	result := bson.NewDocument()

	filter := bson.NewDocument(
		bson.EC.ObjectID("_id", u.Id),
	)

	err := db.Collection("users").FindOne(context.Background(), filter).Decode(result)

	if err != nil {

		return u, err
	}

	firstName := result.Lookup("first_name")

	if firstName != nil {
		u.FirstName = firstName.StringValue()
	}

	lastName := result.Lookup("last_name")

	if lastName != nil {
		u.LastName = lastName.StringValue()
	}

	email := result.Lookup("email")
	if email != nil {
		u.Email = email.StringValue()
	}

	u.Password = ""

	u.Created = result.Lookup("created").DateTime()
	u.Updated = result.Lookup("updated").DateTime()

	return u, err
}

func (u User) validateCreate() (User, error) {

	var err error = nil

	// Email validation
	if u.Email == "" {
		err = errors.New("email is required")
		return u, err
	}

	u.Email = strings.ToLower(u.Email)
	err = helper.ValidateEmail(u.Email)

	if err != nil {
		return u, err
	}

	// trim space
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	// Password validation
	if u.Password == "" {
		err = errors.New("password is required")
		return u, err
	}
	
	if len(u.Password) < 6 {
		err = errors.New("password must be of minimum 6 characters length")
		return u, err
	}

	return u, err
}
