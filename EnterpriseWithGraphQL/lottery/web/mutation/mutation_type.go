package mutation

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/address"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/admin"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/level"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/lottery"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"log"
	"time"

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
	Mutation.AddFieldConfig("sign", &graphql.Field{
		Name:        "sign",
		Type:        admin.TypeAdmin,
		Description: "sign in",
		Args: graphql.FieldConfigArgument{
			"phone": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			phone := p.Args["phone"].(string)
			password := p.Args["password"].(string)
			var params model.LoginParam
			params = model.LoginParam{
				Password: password,
				Phone:    phone,
			}
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return model.CreateAdmin(params)
		},
	})
	Mutation.AddFieldConfig("login", &graphql.Field{
		Name:        "login",
		Type:        admin.TypeAdmin,
		Description: "login the website",
		Args: graphql.FieldConfigArgument{
			"phone": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			phone := p.Args["phone"].(string)
			password := p.Args["password"].(string)
			var params model.LoginParam
			params = model.LoginParam{
				Phone:    phone,
				Password: password,
			}
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return model.Login(params)
		},
	})
}

func init() {

	Mutation.AddFieldConfig("createAddress", &graphql.Field{
		Name:        "crateAddress",
		Type:        address.TypeAddress,
		Description: "create address for one admin",
		Args: graphql.FieldConfigArgument{
			"adminId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"detail": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			adminId := p.Args["adminId"]
			ID := assistance.ToInt64(adminId)
			detail := p.Args["detail"].(string)
			var params model.CreateAddressParams
			params = model.CreateAddressParams{
				AdminId: ID,
				Detail:  detail,
			}
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return model.CreateAddress(params.AdminId, params.Detail)
		},
	})
	Mutation.AddFieldConfig("updateAddress", &graphql.Field{
		Name:        "updateAddress",
		Type:        address.TypeAddress,
		Description: "updateAddress",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"detail": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["id"]
			ID := assistance.ToInt64(id)
			detail := p.Args["detail"].(string)
			return model.UpdateAddress(ID, detail)
		},
	})
}

func init() {
	Mutation.AddFieldConfig("createLottery", &graphql.Field{
		Name:        "createLottery",
		Type:        lottery.TypeLottery,
		Description: "create lottery",
		Args: graphql.FieldConfigArgument{
			"levels": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewList(level.TypeLevel)),
			},
			"deadline": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"winnerConditionId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"class": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"adminId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var params model.CreateLotteryParams
			var deadline time.Time
			log.Println(p.Args)
			if p.Args["deadline"].(string) == "" {
				now := time.Now()
				deadline = now.AddDate(0, 0, 1)
			}

			params = model.CreateLotteryParams{
				Levels:                 p.Args["levels"].([]model.LevelsParams),
				Deadline:               deadline,
				WinnerLotteryCondition: assistance.ToInt64(p.Args["LevelsParams"]),
				LotteryClass:           assistance.ToInt(p.Args["class"]),
				Limit:                  assistance.ToInt(p.Args["limit"]),
				AdminId:                assistance.ToInt64(p.Args["adminId"]),
			}
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return model.CreateLottery(params)
		},
	})
}
