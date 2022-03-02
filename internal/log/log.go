package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
)

// Init the logging engine
func Init() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logLevel := viper.GetString(config.Keys.LogLevel)

	if logLevel != "" {
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			return err
		}
		logrus.SetLevel(level)

		if level == logrus.TraceLevel {
			logrus.SetReportCaller(true)
		}
	}

	return nil
}
