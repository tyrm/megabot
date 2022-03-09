package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/tyrm/megabot/internal/jwt"
)

func (m *Module) loginMutator(params graphql.ResolveParams) (interface{}, error) {
	l := logger.WithField("func", "loginMutator")
	l.Debugf("trying to login")

	// marshall and cast the argument values
	email, _ := params.Args["email"].(string)
	password, _ := params.Args["password"].(string)

	user, err := m.db.ReadUserByEmail(params.Context, email)
	if err != nil {
		l.Errorf("db error: %s", err.Error())
		return nil, err
	}
	if user == nil {
		return nil, errBadLogin
	}

	// check password validity
	if !user.CheckPasswordHash(password) {
		return nil, errBadLogin
	}

	// create jwt
	ts, err := m.jwt.CreateToken(params.Context, user)
	if err != nil {
		l.Debugf("error creating token: %s", err.Error())
		return nil, err
	}

	// save jwt
	err = m.jwt.CreateAuth(params.Context, user.ID, ts)
	if err != nil {
		l.Debugf("error saving token: %s", err.Error())
		return nil, err
	}

	return ts, nil
}

func (m *Module) logoutMutator(params graphql.ResolveParams) (interface{}, error) {
	l := logger.WithField("func", "logoutMutator")
	l.Debugf("trying to logout")

	if params.Context.Value(metadataKey) == nil {
		return nil, errUnauthorized
	}
	metadata := params.Context.Value(metadataKey).(*jwt.AccessDetails)

	err := m.jwt.DeleteTokens(params.Context, metadata)
	if err != nil {
		l.Tracef("can't delete tokens: %s", err.Error())
		return nil, err
	}

	return success{Success: true}, nil
}

func (m *Module) refreshAccessTokenMutator(params graphql.ResolveParams) (interface{}, error) {
	l := logger.WithField("func", "refreshAccessTokenMutator")
	l.Debugf("trying to refresh token")

	// marshall and cast the argument values
	refreshToken, _ := params.Args["refreshToken"].(string)

	return m.jwt.RefreshAccessToken(params.Context, refreshToken)
}
