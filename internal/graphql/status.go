package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
)

type status struct {
	Version string `json:"version"`
}

func (m *Module) statusQuery(params graphql.ResolveParams) (interface{}, error) {
	logrus.Debugf("trying to get status")

	newStatus := status{
		Version: viper.GetString(config.Keys.SoftwareVersion),
	}

	return newStatus, nil
}
