package mutation

import (
	"net/http"

	"github.com/graphql-go/graphql"
)

type Method struct {
	Name      string `json:"name"`
	Operation string `json:"operation"`
}

var defaultMethod = Method{
	Name:      http.MethodPost,
	Operation: "mutation",
}

func makeMethod(name string) *Method {
	return &Method{
		Name:      name,
		Operation: "mutation",
	}
}

var TypeMethod = graphql.NewObject(graphql.ObjectConfig{
	Name: "Method",
	Fields: graphql.Fields{
		"method": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if method, ok := p.Source.(*Method); ok {
					return method.Name, nil
				}
				return nil, nil
			},
		},
		"operation": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if method, ok := p.Source.(*Method); ok {
					return method.Operation, nil
				}
				return nil, nil
			},
		},
	},
})
