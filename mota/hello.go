package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Hello struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Hello",
	Fields: graphql.Fields{
		"from": &graphql.Field{
			Type: graphql.String,
		},
		"to": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: helloType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				myHello := Hello{
					From:    "me",
					To:      "me",
					Message: "hello",
				}
				return myHello, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
