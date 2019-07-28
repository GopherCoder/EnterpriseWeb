package admin

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"

	"github.com/graphql-go/graphql"
)

var TypeAdmin = graphql.NewObject(graphql.ObjectConfig{
	Name: "Admin",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if admin, ok := p.Source.(*model.AdminSerializer); ok == true {
					return admin.Id, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if admin, ok := p.Source.(*model.AdminSerializer); ok == true {
					return admin.Name, nil
				}
				return nil, nil
			},
		},
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if admin, ok := p.Source.(*model.AdminSerializer); ok == true {
					return admin.Token, nil
				}
				return nil, nil
			},
		},
		"phone": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if admin, ok := p.Source.(*model.AdminSerializer); ok == true {
					return admin.Phone, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {

}
