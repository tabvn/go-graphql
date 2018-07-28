package model

import (
	"time"
	"github.com/graphql-go/graphql"
	"errors"
	"go-graphql/db"
	"database/sql"
	"github.com/satori/go.uuid"
	)

type Token struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
	Created int64  `json:"created"`
}

var TokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Token",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"created": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func (token *Token) Create() (*Token, error) {


	if token.Token == "" {
		token.Token = uuid.Must(uuid.NewV4()).String()
	}
	query := `INSERT INTO tokens (user_id, token, created) VALUES (?, ?, ?)`
	currentTime := time.Now()
	token.Created = currentTime.Unix()

	result, err := db.DB.Insert(query, token.UserId, token.Token, token.Created)

	if err != nil {
		return nil, err
	}

	token.Id = result

	return token, err
}

func scanToken(s db.RowScanner) (*Token, error) {
	var (
		id      int64
		userId  sql.NullInt64
		token   sql.NullString
		created sql.NullInt64
	)
	if err := s.Scan(&id, &userId, &token, &created); err != nil {
		return nil, err
	}

	t := &Token{
		Id:      id,
		UserId:  userId.Int64,
		Token:   token.String,
		Created: created.Int64,
	}

	return t, nil
}

func (token *Token) Load() (*Token, error) {

	row, err := db.DB.Get("tokens", token.Id)

	if err != nil {
		return nil, err
	}

	t, err := scanToken(row)

	if t == nil {
		return nil, errors.New("token not found")
	}

	return t, err
}
