package graphql

import "github.com/graphql-go/graphql"

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
