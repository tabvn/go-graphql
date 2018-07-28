package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"go-graphql/schema"
	"go-graphql/db"
	"errors"
	"go-graphql/dev"
	"go-graphql/config"
)

type params struct {
	Query         string      `json:"query"`
	OperationName string      `json:"operationName,omitempty"`
	Variables     interface{} `json:"variables,omitempty"`
}

func getBodyFromRequest(r *http.Request) (string, error) {

	var body string

	if r.Method == "GET" {
		body = r.URL.Query().Get("query")

	}

	if r.Method == "POST" {
		p := &params{
			Variables: nil,
		}

		if err := json.NewDecoder(r.Body).Decode(p); err != nil {
			return body, err
		}
		body = p.Query

	}

	return body, nil
}

func Setup() {

	_, err := db.InitDatabase()
	if err != nil {
		panic(errors.New("can not connect to database"))
	}

}

func main() {

	Setup()

	// GraphIQL development router handler
	if !config.Production {
		http.HandleFunc("/", dev.ServeGraphiQL)
	}

	// Router api graphQL handler
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		body, error := getBodyFromRequest(r)
		if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
		}
		result := schema.ExecuteQuery(body, schema.Schema)

		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 3001")
	http.ListenAndServe(":3001", nil)
}
