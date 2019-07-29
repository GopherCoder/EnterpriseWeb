package mutation

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/log"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/address"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/admin"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/level"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/lottery"
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
			var params admin.LoginParam
			params = admin.LoginParam{
				Password: password,
				Phone:    phone,
			}
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return admin.CreateAdmin(params)
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
			var params admin.LoginParam
			params = admin.LoginParam{
				Phone:    phone,
				Password: password,
			}
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return admin.Login(params)
		},
	})
	Mutation.AddFieldConfig("updateAdmin", &graphql.Field{
		Type: admin.TypeAdmin,
		Args: graphql.FieldConfigArgument{
			"adminId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var params admin.UpdateAdminParams
			params.AdminId = assistance.ToInt64(p.Args["adminId"])
			params.Name = p.Args["name"].(string)
			return admin.UpdateAdmin(params)
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
			log_for_lottery.Println(p.Args)
			adminId := p.Args["adminId"]
			ID := assistance.ToInt64(adminId)
			detail := p.Args["detail"].(string)
			var params address.CreateAddressParams
			params = address.CreateAddressParams{
				AdminId: ID,
				Detail:  detail,
			}
			log_for_lottery.Println(params)
			if err := params.Valid(); err != nil {
				return nil, err
			}
			return address.CreateAddress(params)
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
			var params address.UpdateAddressParams
			params = address.UpdateAddressParams{
				Id:     ID,
				Detail: detail,
			}
			return address.UpdateAddress(params)
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
				Type: graphql.NewList(level.TypeLevelsInput),
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
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var params lottery.CreateLotteryParams
			log_for_lottery.Println(p.Args)

			var deadline time.Time
			if p.Args["deadline"] == nil {
				now := time.Now()
				deadline = now.AddDate(0, 0, 1)
			} else {
				deadline = p.Args["deadline"].(time.Time)
			}

			var levels []lottery.LevelsParams
			for _, k := range p.Args["levels"].([]interface{}) {
				i := k.(map[string]interface{})
				var one lottery.LevelsParams
				one = lottery.LevelsParams{
					Name: i["name"].(string),
				}
				if i["imageURL"] != nil {
					one.ImageURL = i["imageURL"].(string)
				}
				if i["number"] != nil {
					one.Number = i["number"].(int)
				}
				if i["class"] != nil {
					one.Class = i["class"].(int)
				}
				levels = append(levels, one)

			}

			params = lottery.CreateLotteryParams{
				Levels:                 levels,
				Deadline:               deadline,
				WinnerLotteryCondition: int64(p.Args["winnerConditionId"].(int)),
				LotteryClass:           p.Args["class"].(int),
				Limit:                  p.Args["limit"].(int),
				AdminId:                assistance.ToInt64(p.Args["adminId"]),
			}
			log_for_lottery.Println(params)
			if err := params.Valid(); err != nil {
				return nil, err
			}

			return lottery.CreateLottery(params)
		},
	})
	Mutation.AddFieldConfig("takePartIn", &graphql.Field{
		Name: "takePartIn",
		Type: lottery.TypeLottery,
		Args: graphql.FieldConfigArgument{
			"adminId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"lotteryId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var params lottery.TakePartInParams
			params.AdminId = assistance.ToInt64(p.Args["adminId"])
			params.LotteryId = assistance.ToInt64(p.Args["lotteryId"])
			return lottery.TakePartInOneLottery(params)
		},
	})
}
