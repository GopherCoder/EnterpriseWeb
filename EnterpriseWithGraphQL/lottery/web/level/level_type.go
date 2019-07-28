package level

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

var TypeLevel = graphql.NewObject(graphql.ObjectConfig{
	Name: "level",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.Level); ok {
					return level.Id, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.Level); ok {
					return level.Name, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"image_url": &graphql.Field{
			Name: "image_url",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.Level); ok {
					return level.ImageURL, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"number": &graphql.Field{
			Name: "number",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.Level); ok {
					return level.Number, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"class": &graphql.Field{
			Name: "class",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.Level); ok {
					return level.Class, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})
