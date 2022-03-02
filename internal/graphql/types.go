package graphql

import "github.com/graphql-go/graphql"

type success struct {
	Success bool `json:"success"`
}

// input types
var jwtTokensType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JwtToken",
	Fields: graphql.Fields{
		"accessToken": &graphql.Field{
			Type: graphql.String,
		},
		"refreshToken": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var statusType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Status",
	Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var successType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Success",
	Fields: graphql.Fields{
		"success": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
