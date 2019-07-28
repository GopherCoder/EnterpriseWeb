package query

import (
	"net/http"

	"github.com/graphql-go/graphql"
)

type Ping struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

var defaultPing = Ping{
	Data: "pong",
	Code: http.StatusOK,
}

func makePing(data string) Ping {
	return Ping{
		Data: data,
		Code: http.StatusOK,
	}
}

var PingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ping",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if ping, ok := p.Source.(*Ping); ok == true {
					return ping.Data, nil
				}
				return nil, nil
			},
		},
		"code": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if ping, ok := p.Source.(*Ping); ok == true {
					return ping.Code, nil
				}
				return nil, nil
			},
		},
	},
})
