package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"go-graphql/helper"
	"errors"
	"go-graphql/db"
)

type User struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{

			"id": &graphql.Field{
				Type: graphql.ID,
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

	// generate password
	password, e := HashPassword(u.Password)
	u.Password = password

	if e != nil {
		return u, e
	}

	query := `INSERT INTO users (first_name, last_name, email, password, created, updated) VALUES (?, ?, ?, ?, ?, ?)`
	currentTime := time.Now()
	u.Created = currentTime.Unix()
	u.Updated = currentTime.Unix()

	result, err := db.DB.Insert(query, u.FirstName, u.LastName, u.Email, u.Password, u.Created, u.Updated)

	if err != nil {
		return u, err
	}

	u.Id = result
	u.Password = ""

	return u, err
}

func (u User) Update() (User, error) {

	currentTime := time.Now()
	u.Updated = currentTime.Unix()

	if u.Password == "" {
		query := `UPDATE users SET first_name=?, last_name=?, email=?, updated=? WHERE id = ?`
		_, err := db.DB.Update(query, u.FirstName, u.LastName, u.Email, u.Updated, u.Id)

		if err != nil {
			return u, err
		}
	} else {
		query := `UPDATE users SET first_name=?, last_name=?, email=?, password=?, updated=? WHERE id = ?`
		password, err := HashPassword(u.Password)
		if err != nil {
			return u, err
		}
		_, updateErr := db.DB.Update(query, u.FirstName, u.LastName, u.Email, password, u.Updated, u.Id)

		if updateErr != nil {
			return u, err
		}
	}

	u.Password = ""

	return u, nil
}

func (u User) Load() (User, error) {

	return u, nil
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

	//@todo validate email in database if exist.

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
