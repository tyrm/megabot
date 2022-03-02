package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

func (m *Module) loginMutator(params graphql.ResolveParams) (interface{}, error) {
	logrus.Debugf("trying to login")

	// marshall and cast the argument values
	email, _ := params.Args["email"].(string)
	password, _ := params.Args["password"].(string)

	user, err := m.db.ReadUserByEmail(params.Context, email)
	if err != nil {
		logrus.Errorf("db error: %s", err.Error())
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
		logrus.Debugf("error creating token: %s", err.Error())
		return nil, err
	}

	// save jwt
	err = m.jwt.CreateAuth(params.Context, user.ID, ts)
	if err != nil {
		logrus.Debugf("error saving token: %s", err.Error())
		return nil, err
	}

	return ts, nil
}
