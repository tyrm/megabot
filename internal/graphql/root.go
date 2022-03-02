package graphql

import "github.com/graphql-go/graphql"

func (m *Module) statusField() *graphql.Field {
	return &graphql.Field{
		Type:        statusType,
		Description: "get system status",
		Resolve:     m.statusQuery,
	}
}

func (m *Module) rootQueryUnauthorized() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"status": m.statusField(),
		},
	})
}

func (m *Module) rootMutationUnauthorized() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type:        jwtTokensType,
				Description: "Login to system",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: m.loginMutator,
			},
		},
	})
}
