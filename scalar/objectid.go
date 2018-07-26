package scalar

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type ID struct {
	value objectid.ObjectID
}

func (id *ID) String() objectid.ObjectID {
	return id.value
}

func NewId(v string) *ID {

	id, err := objectid.FromHex(v)
	if err != nil {
		return nil
	}
	return &ID{value: id}

}

var ObjectIdType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ObjectId",
	Description: "The `ObjectIdType` scalar type represents an ID Object.",
	// Serialize serializes `CustomID` to string.
	Serialize: func(value interface{}) interface{} {

		switch value := value.(type) {

		case objectid.ObjectID:
			return value.Hex()

		case ID:
			return value.String()
		case *ID:
			v := *value
			return v.String()
		default:
			return nil
		}
	},

	// ParseValue parses GraphQL variables from `string` to `ObjectId`.
	ParseValue: func(value interface{}) interface{} {

		switch value := value.(type) {

		case string:

			return NewId(value)
		case *string:
			return NewId(*value)
		default:
			return nil
		}
	},
	// ParseLiteral parses GraphQL AST value to `CustomID`.
	ParseLiteral: func(valueAST ast.Value) interface{} {

		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return NewId(valueAST.Value)
		default:
			return nil
		}
	},
})
