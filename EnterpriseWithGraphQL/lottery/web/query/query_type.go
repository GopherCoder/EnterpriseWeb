package query

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/address"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/admin"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type:        PingType,
			Description: "health check",
			Args: graphql.FieldConfigArgument{
				"data": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "data to ping",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				data := p.Args["data"]
				if data != nil {
					ping := makePing(data.(string))
					return &ping, nil
				}
				return &defaultPing, nil

			},
		},
	},
})

func init() {
	Query.AddFieldConfig("admin", &graphql.Field{
		Type: admin.TypeAdmin,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "id of admin",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["id"]
			ID := assistance.ToInt64(id)
			if ID == 0 {
				return model.DefaultAdmin()
			}
			result, err := model.GetAdmin(int64(ID))
			if err != nil {
				return nil, err
			} else {
				return result, err
			}
		},
	})
}

func init() {

	Query.AddFieldConfig("address", &graphql.Field{
		Type: graphql.NewList(address.TypeAddress),
		Args: graphql.FieldConfigArgument{
			"adminId": &graphql.ArgumentConfig{
				Description: "id of admin",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["adminId"]
			ID := assistance.ToInt64(id)
			return model.GetAddresses(ID)
		},
	})
}

func init() {
	//Query.AddFieldConfig("lotteries", &graphql.Field{
	//	Type: graphql.NewNonNull(graphql.NewList(lottery.TypeLottery)),
	//	Args: graphql.FieldConfigArgument{
	//		"ownerId": &graphql.ArgumentConfig{
	//			Description: "id of admin",
	//			Type:        graphql.NewNonNull(graphql.ID),
	//		},
	//	},
	//	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
	//		id := p.Args["ownerId"]
	//		ID := assistance.ToInt64(id)
	//		return model.ListLottery(ID)
	//	},
	//})
	//Query.AddFieldConfig("lottery", &graphql.Field{
	//	Type: lottery.TypeLottery,
	//	Args: graphql.FieldConfigArgument{
	//		"id": &graphql.ArgumentConfig{
	//			Description: "id of lottery",
	//			Type:        graphql.NewNonNull(graphql.ID),
	//		},
	//	},
	//	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
	//		id := p.Args["id"]
	//		return model.OneLottery(id.(int64))
	//	},
	//})
	//Query.AddFieldConfig("involvements", &graphql.Field{
	//	Type: graphql.NewNonNull(graphql.NewList(lottery.TypeLottery)),
	//	Args: graphql.FieldConfigArgument{
	//		"adminId": &graphql.ArgumentConfig{
	//			Description: "id of admin",
	//			Type:        graphql.NewNonNull(graphql.ID),
	//		},
	//	},
	//	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
	//		id := p.Args["adminId"]
	//		return model.InvolvementsLottery(id.(int64))
	//	},
	//})
}
