package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

func (m *Module) schema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    m.rootQueryUnauthorized(),
		Mutation: m.rootMutationUnauthorized(),
	})
	if err != nil {
		logrus.Errorf("can't create schema: %s", err.Error())
	}
	return schema
}

func (m *Module) schemaUnauthorized() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    m.rootQueryUnauthorized(),
		Mutation: m.rootMutationUnauthorized(),
	})
	if err != nil {
		logrus.Errorf("can't create schema: %s", err.Error())
	}
	return schema
}
