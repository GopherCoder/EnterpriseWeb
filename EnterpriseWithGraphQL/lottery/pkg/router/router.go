package router

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/mutation"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/query"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		opts := NewRequestOptions(request)
		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
		})

		message := fmt.Sprintf("%s | %s | %s | %s", request.Method, request.Host, request.RequestURI, time.Now().Format(time.RFC3339))
		log.Println(string(opts.Query))
		log.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", message))

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(result)
	}
}

func CollectionOfRouter() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.Query,
		Mutation: mutation.Mutation,
	})
	if err != nil {
		log.Fatal("New Schema Fail :", err.Error())
		return
	}
	h := handler(schema)
	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8999", nil))
}
