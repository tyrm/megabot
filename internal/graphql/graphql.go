package graphql

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/jwt"
	"github.com/tyrm/megabot/internal/web"
	"net/http"
)

const PathGraphQL = "/graphql"

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// Module implements the web module interface
type Module struct {
	db  db.DB
	jwt *jwt.Module
}

// New returns a new auth module
func New(db db.DB, jwt *jwt.Module) web.Module {
	return &Module{
		db:  db,
		jwt: jwt,
	}
}

func (m *Module) Route(s web.Server) error {
	s.HandleFunc(PathGraphQL, m.graphqlHandler).Methods("POST")
	return nil
}

func (m *Module) graphqlHandler(w http.ResponseWriter, r *http.Request) {
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

	logrus.Tracef("query: %s", query)
	logrus.Tracef("operation: %s", operation)
	logrus.Tracef("variables: %v", variables)

	ctx := r.Context()

	// check auth
	var err error
	//metadata, err := s.extractTokenMetadata(r)

	// do
	var result *graphql.Result
	if err == nil {
		// authorized
		//ctx = context.WithValue(ctx, metadataKey, metadata)
		result = graphql.Do(graphql.Params{
			Context:        ctx,
			Schema:         m.schema(),
			RequestString:  query,
			VariableValues: variables,
			OperationName:  operation,
		})
	} else {
		// unauthorized
		result = graphql.Do(graphql.Params{
			Context:        ctx,
			Schema:         m.schemaUnauthorized(),
			RequestString:  query,
			VariableValues: variables,
			OperationName:  operation,
		})
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Printf("could not write result to response: %s", err)
	}
}
