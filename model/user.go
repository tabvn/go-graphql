package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"go-graphql/helper"
	"errors"
	"go-graphql/db"
	"database/sql"
	"fmt"
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
				Type: graphql.Int,
			},
			"updated": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var LoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "login",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"created": &graphql.Field{
				Type: graphql.Int,
			},
			"user": &graphql.Field{
				Type: UserType,
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

func (u *User) Create() (*User, error) {

	u, validateError := u.validateCreate()

	if validateError != nil {

		return nil, validateError
	}

	// generate password
	password, e := HashPassword(u.Password)
	u.Password = password

	if e != nil {
		return nil, e
	}

	query := `INSERT INTO users (first_name, last_name, email, password, created, updated) VALUES (?, ?, ?, ?, ?, ?)`
	currentTime := time.Now()
	u.Created = currentTime.Unix()
	u.Updated = currentTime.Unix()

	result, err := db.DB.Insert(query, u.FirstName, u.LastName, u.Email, u.Password, u.Created, u.Updated)

	if err != nil {
		return nil, err
	}

	u.Id = result
	u.Password = ""

	return u, err
}

func (u *User) Update() (*User, error) {

	currentTime := time.Now()
	u.Updated = currentTime.Unix()

	if u.Password == "" {
		query := `UPDATE users SET first_name=?, last_name=?, email=?, updated=? WHERE id = ?`
		_, err := db.DB.Update(query, u.FirstName, u.LastName, u.Email, u.Updated, u.Id)

		if err != nil {
			return nil, err
		}
	} else {
		query := `UPDATE users SET first_name=?, last_name=?, email=?, password=?, updated=? WHERE id = ?`
		password, err := HashPassword(u.Password)
		if err != nil {
			return nil, err
		}
		_, updateErr := db.DB.Update(query, u.FirstName, u.LastName, u.Email, password, u.Updated, u.Id)

		if updateErr != nil {
			return nil, err
		}
	}

	u.Password = ""

	return u, nil
}

// scanBook reads a book from a sql.Row or sql.Rows
func scanUser(s db.RowScanner) (*User, error) {
	var (
		id        int64
		firstName sql.NullString
		lastName  sql.NullString
		email     sql.NullString
		password  sql.NullString
		created   sql.NullInt64
		updated   sql.NullInt64
	)
	if err := s.Scan(&id, &firstName, &lastName, &email, &password,
		&created, &updated); err != nil {
		return nil, err
	}

	user := &User{
		Id:        id,
		FirstName: firstName.String,
		LastName:  lastName.String,
		Email:     email.String,
		Password:  password.String,
		Created:   created.Int64,
		Updated:   updated.Int64,
	}

	return user, nil
}

func (u *User) Load() (*User, error) {

	row, err := db.DB.Get("users", u.Id)
	if err != nil {
		return nil, err
	}

	user, err := scanUser(row)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (u *User) validateCreate() (*User, error) {

	var err error = nil

	// Email validation
	if u.Email == "" {
		err = errors.New("email is required")
		return nil, err
	}

	u.Email = strings.ToLower(u.Email)
	err = helper.ValidateEmail(u.Email)

	if err != nil {
		return nil, err
	}

	count, countErr := db.DB.Count("SELECT COUNT(*) FROM users WHERE email=?", u.Email)

	if countErr != nil {
		return nil, errors.New("unable validate email")
	}
	if count > 0 {
		return nil, errors.New("email already exist")
	}

	// trim space
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	// Password validation
	if u.Password == "" {
		err = errors.New("password is required")
		return nil, err
	}

	if len(u.Password) < 6 {
		err = errors.New("password must be of minimum 6 characters length")
		return nil, err
	}

	return u, err
}

func LoginUser(email string, password string) (*Token, *User, error) {

	row := db.DB.QueryRow("SELECT * FROM users WHERE email=?", email)

	user, err := scanUser(row)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {

		return nil, nil, errors.New("login failure")
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, nil, errors.New("login failure")
	}

	fmt.Println("Login", user, err)

	return nil, user, nil

}
