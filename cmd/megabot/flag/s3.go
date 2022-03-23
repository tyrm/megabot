package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/internal/config"
)

// S3 adds flags that are common to s3.
func S3(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.S3Endpoint, values.S3Endpoint, usage.S3Endpoint)
	cmd.PersistentFlags().String(config.Keys.S3Region, values.S3Region, usage.S3Region)
	cmd.PersistentFlags().String(config.Keys.S3AccessKeyID, values.S3AccessKeyID, usage.S3AccessKeyID)
	cmd.PersistentFlags().String(config.Keys.S3SecretAccessKey, values.S3SecretAccessKey, usage.S3SecretAccessKey)
	cmd.PersistentFlags().Bool(config.Keys.S3UseSSL, values.S3UseSSL, usage.S3UseSSL)
	cmd.PersistentFlags().String(config.Keys.S3Bucket, values.S3Bucket, usage.S3Bucket)
	cmd.PersistentFlags().Duration(config.Keys.S3PresignedURLExpiration, values.S3PresignedURLExpiration, usage.S3PresignedURLExpiration)
}
