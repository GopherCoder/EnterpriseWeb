package winner_lottery

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

var TypeEnumWinnerCondition = graphql.NewEnum(graphql.EnumConfig{
	Name:        "winnerCondition",
	Description: "中奖条件",
	Values: graphql.EnumValueConfigMap{
		"TIMELEVEL": &graphql.EnumValueConfig{
			Value:       model.TIMELEVEL,
			Description: model.WinnerCondition[model.TIMELEVEL],
		},
		"PERSONLEVEL": &graphql.EnumValueConfig{
			Value:       model.PERSONLEVEL,
			Description: model.WinnerCondition[model.PERSONLEVEL],
		},
		"NOWLEVEL": &graphql.EnumValueConfig{
			Value:       model.NOWLEVEL,
			Description: model.WinnerCondition[model.NOWLEVEL],
		},
	},
})

var TypeWinnerLottery = graphql.NewObject(graphql.ObjectConfig{
	Name: "winnerLottery",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if w, ok := p.Source.(*model.WinnerLotterySerializer); ok {
					return w.Id, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if w, ok := p.Source.(*model.WinnerLotterySerializer); ok {
					return w.Description, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"class": &graphql.Field{
			Type: TypeEnumWinnerCondition,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if w, ok := p.Source.(*model.WinnerLotterySerializer); ok {
					return w.Class, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"classString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if w, ok := p.Source.(*model.WinnerLotterySerializer); ok {
					return w.ClassString, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})
