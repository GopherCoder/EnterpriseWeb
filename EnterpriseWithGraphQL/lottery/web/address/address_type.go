package address

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

var TypeAddress = graphql.NewObject(graphql.ObjectConfig{
	Name: "Address",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if address, ok := p.Source.(*model.Address); ok {
					return address.Id, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"detail": &graphql.Field{
			Name: "detail",
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if address, ok := p.Source.(*model.Address); ok {
					return address.Detail, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})
