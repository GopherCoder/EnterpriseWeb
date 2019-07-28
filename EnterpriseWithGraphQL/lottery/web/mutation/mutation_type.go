package mutation

import (
	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"method": &graphql.Field{
			Type:        TypeMethod,
			Description: "return all method",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				name, _ := p.Args["name"]
				if name == nil {
					return &defaultMethod, nil
				}
				return makeMethod(name.(string)), nil
			},
		},
	},
})

func init() {
	//Mutation.AddFieldConfig("sign", &graphql.Field{})
}
