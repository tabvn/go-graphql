package schema

import (
	"github.com/graphql-go/graphql"
	"go-graphql/query"
	"go-graphql/mutation"
	"fmt"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    query.Query,
		Mutation: mutation.Mutation,
	},
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}
