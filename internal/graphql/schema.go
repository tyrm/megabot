package graphql

import (
	"github.com/graphql-go/graphql"
)

func (m *Module) schema() graphql.Schema {
	l := logger.WithField("func", "schema")
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    m.rootQuery(),
		Mutation: m.rootMutation(),
	})
	if err != nil {
		l.Errorf("can't create schema: %s", err.Error())
	}
	return schema
}

func (m *Module) schemaUnauthorized() graphql.Schema {
	l := logger.WithField("func", "schemaUnauthorized")
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    m.rootQueryUnauthorized(),
		Mutation: m.rootMutationUnauthorized(),
	})
	if err != nil {
		l.Errorf("can't create schema: %s", err.Error())
	}
	return schema
}
