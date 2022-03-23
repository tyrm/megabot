package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"time"
)

// Module is an s3 capable filestore
type Module struct {
	mc                     *minio.Client
	bucket                 string
	presignedURLExpiration time.Duration
}

// New returns a new S3 capable module for filestore.Filestore
func New() (*Module, error) {
	l := logger.WithField("func", "New")

	endpoint := viper.GetString(config.Keys.S3Endpoint)
	accessKeyID := viper.GetString(config.Keys.S3AccessKeyID)
	secretAccessKey := viper.GetString(config.Keys.S3SecretAccessKey)
	useSSL := viper.GetBool(config.Keys.S3UseSSL)

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		l.Errorf("creating minio client: %s", err.Error())
		return nil, err
	}

	return &Module{
		mc:                     mc,
		bucket:                 viper.GetString(config.Keys.S3Bucket),
		presignedURLExpiration: viper.GetDuration(config.Keys.S3PresignedURLExpiration),
	}, nil
}
