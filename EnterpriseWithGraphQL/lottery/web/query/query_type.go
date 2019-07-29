package query

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/log"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/address"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/admin"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/lottery"
	"log"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type:        TypePing,
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
				return admin.DefaultAdmin()
			}
			result, err := admin.GetAdmin(int64(ID))
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
			"orderBy": &graphql.ArgumentConfig{
				Description: "order by id desc",
				Type:        graphql.String,
			},
			"limit": &graphql.ArgumentConfig{
				Description: "limit",
				Type:        graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Description: "offset",
				Type:        graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			log_for_lottery.Println(p.Args)
			var params address.GetAddressParams
			orderBy := p.Args["orderBy"]
			if orderBy == nil {
				orderBy = "created_at"
			}
			params.OrderBy = orderBy.(string)
			offset := p.Args["offset"]
			if offset == nil {
				offset = 0
			}
			params.Offset = offset.(int)
			limit := p.Args["limit"]
			if limit != nil {
				params.Limit = limit.(int)
			}
			params.AdminId = assistance.ToInt64(p.Args["adminId"])
			log.Println(params)
			return address.GetAddresses(params)
		},
	})
}

func init() {
	Query.AddFieldConfig("lotteries", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(lottery.TypeLottery)),
		Args: graphql.FieldConfigArgument{
			"ownerId": &graphql.ArgumentConfig{
				Description: "id of admin",
				Type:        graphql.NewNonNull(graphql.ID),
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["ownerId"]
			ID := assistance.ToInt64(id)
			var params lottery.ListLotteryParams
			params.OwnerId = ID
			if p.Args["orderBy"] != nil {
				params.OrderBy = p.Args["orderBy"].(string)
			} else {
				params.OrderBy = "created_at"
			}
			if p.Args["limit"] != nil {
				params.Limit = p.Args["limit"].(int)
			} else {
				params.Limit = 10
			}
			if p.Args["offset"] != nil {
				params.Offset = p.Args["offset"].(int)
			} else {
				params.Offset = 0
			}
			return lottery.ListLottery(params)
		},
	})
	Query.AddFieldConfig("lottery", &graphql.Field{
		Type: lottery.TypeLottery,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "id of lottery",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["id"]
			return lottery.OneLottery(assistance.ToInt64(id))
		},
	})
	Query.AddFieldConfig("involvements", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(lottery.TypeLottery)),
		Args: graphql.FieldConfigArgument{
			"adminId": &graphql.ArgumentConfig{
				Description: "id of admin",
				Type:        graphql.NewNonNull(graphql.ID),
			},
			"orderBy": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["adminId"]
			var params lottery.InvolvementParams
			params.AdminId = assistance.ToInt64(id)
			if p.Args["orderBy"] != nil {
				params.OrderBy = p.Args["orderBy"].(string)
			}
			if p.Args["limit"] != nil {
				params.Limit = p.Args["limit"].(int)
			} else {
				params.Limit = 10
			}
			return lottery.InvolvementsLottery(params)
		},
	})
}
