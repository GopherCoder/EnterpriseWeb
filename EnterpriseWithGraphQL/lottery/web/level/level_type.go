package level

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

var TypeLevelEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "levelEnum",
	Description: "level of enum",
	Values: graphql.EnumValueConfigMap{
		"FIRST": &graphql.EnumValueConfig{
			Value:       model.FIRST,
			Description: model.Prize[model.FIRST],
		},
		"SECOND": &graphql.EnumValueConfig{
			Value:       model.SECOND,
			Description: model.Prize[model.SECOND],
		},
		"THIRD": &graphql.EnumValueConfig{
			Value:       model.THIRD,
			Description: model.Prize[model.THIRD],
		},
		"FOURTH": &graphql.EnumValueConfig{
			Value:       model.FOURTH,
			Description: model.Prize[model.FOURTH],
		},
		"FIFTH": &graphql.EnumValueConfig{
			Value:       model.FIFTH,
			Description: model.Prize[model.FIFTH],
		},
		"SIXTH": &graphql.EnumValueConfig{
			Value:       model.SIXTH,
			Description: model.Prize[model.SIXTH],
		},
	},
})
var TypeLevelsInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "levelsInput",
	Description: "input params",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"imageURL": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"number": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"class": &graphql.InputObjectFieldConfig{
			Type: TypeLevelEnum,
		},
	},
})
var TypeLevel = graphql.NewObject(graphql.ObjectConfig{
	Name: "level",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.LevelSerializer); ok {
					return level.Id, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.LevelSerializer); ok {
					return level.Name, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"imageURL": &graphql.Field{
			Name: "image_url",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.LevelSerializer); ok {
					return level.ImageURL, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"number": &graphql.Field{
			Name: "number",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.LevelSerializer); ok {
					return level.Number, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"class": &graphql.Field{
			Name: "class",
			Type: TypeLevelEnum,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.LevelSerializer); ok {
					//fmt.Println(fmt.Sprintf("%#v", level))
					return level.Class, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"classString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if level, ok := p.Source.(*model.LevelSerializer); ok {
					return level.ClassString, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})
