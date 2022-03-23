package flag

import "github.com/tyrm/megabot/internal/config"

var usage = config.KeyNames{
	ConfigPath: "Path to a file containing megabot configuration. Values set in this file will be overwritten by values set as env vars or arguments",
	LogLevel:   "Log level to run at: [trace, debug, info, warn, fatal]",

	// application
	ApplicationName: "Name of the application, used in various places internally",

	// database
	DbType:         "Database type: eg., postgres",
	DbAddress:      "Database ipv4 address, hostname, or filename",
	DbPort:         "Database port",
	DbUser:         "Database username",
	DbPassword:     "Database password",
	DbDatabase:     "Database name",
	DbTLSMode:      "Database tls mode",
	DbTLSCACert:    "Path to CA cert for db tls connection",
	DbLoadTestData: "Should test data be loaded into the database",

	// s3
	S3Endpoint:               "S3 Endpoint",
	S3Region:                 "S3 Region",
	S3AccessKeyID:            "S3 Access ID",
	S3SecretAccessKey:        "S3 Secret Key",
	S3UseSSL:                 "S3 Use SSL",
	S3Bucket:                 "S3 Bucket",
	S3PresignedURLExpiration: "S3 Presigned Expiration Duration",

	// server
	ServerExternalHostname: "The external hostname used by the server",
	ServerHTTP2:            "Enable HTTP2 Server",
	ServerHTTP3:            "Enable HTTP3 Server",
	ServerMinifyHTML:       "Should the server minify html documents before sending",
	ServerRoles:            "Server roles that should be started: [graphql, webapp]",
	ServerTLSCertPath:      "TLS server cert path",
	ServerTLSKeyPath:       "TLS server key path",

	// user
	UserEmail:    "User email address",
	UserGroups:   "User groups",
	UserPassword: "User password",
}
