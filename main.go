package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"go-graphql/schema"
)

func getBodyFromRequest(r *http.Request) (string, error) {

	var body string

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if r.Method == "GET" {
		body = r.URL.Query().Get("query")
	}
	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			return body, err
		}
		body = params.Query

	}

	return body, nil
}

func main() {

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
