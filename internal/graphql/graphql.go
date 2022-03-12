package graphql

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/jwt"
	"github.com/tyrm/megabot/internal/web"
	"net/http"
)

const pathGraphQL = "/graphql"

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// Module contains a graphql module for the web server. Implements web.Module
type Module struct {
	db  db.DB
	jwt *jwt.Module
}

// New returns a new graphql module
func New(db db.DB, jwt *jwt.Module) web.Module {
	return &Module{
		db:  db,
		jwt: jwt,
	}
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleGraphQL
}

// Route attaches routes to the web server
func (m *Module) Route(s web.Server) error {
	s.HandleFunc(pathGraphQL, m.graphqlHandler).Methods("POST")
	return nil
}

func (m *Module) graphqlHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "graphqlHandler")
	var p map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(400)
		return
	}

	var query string
	var operation string
	var variables map[string]interface{}

	if val, ok := p["query"].(string); ok {
		query = val
	}
	if val, ok := p["operation"].(string); ok {
		operation = val
	}
	if val, ok := p["variables"].(map[string]interface{}); ok {
		variables = val
	}

	ctx := r.Context()

	// check auth
	var err error
	metadata, err := m.jwt.ExtractTokenMetadata(r)
	if err != nil {
		l.Debugf("extract token metadata error: %s", err.Error())
	}

	// do
	var result *graphql.Result
	if err == nil {
		// authorized
		l.Tracef("authorzed query: %s", query)
		ctx = context.WithValue(ctx, metadataKey, metadata)
		result = graphql.Do(graphql.Params{
			Context:        ctx,
			Schema:         m.schema(),
			RequestString:  query,
			VariableValues: variables,
			OperationName:  operation,
		})
	} else {
		// unauthorized
		l.Tracef("unauthorized query: %s", query)
		result = graphql.Do(graphql.Params{
			Context:        ctx,
			Schema:         m.schemaUnauthorized(),
			RequestString:  query,
			VariableValues: variables,
			OperationName:  operation,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		l.Errorf("could not write result to response: %s", err)
	}
}
