package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName string
	SoftwareVersion string
	TokenSalt       string

	// database
	DbType          string
	DbAddress       string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbDatabase      string
	DbTLSMode       string
	DbTLSCACert     string
	DbLoadTestData  string
	DbEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// auth
	AccessExpiration  string
	AccessSecret      string
	RefreshExpiration string
	RefreshSecret     string

	// s3
	S3Endpoint               string
	S3Region                 string
	S3AccessKeyID            string
	S3SecretAccessKey        string
	S3UseSSL                 string
	S3Bucket                 string
	S3PresignedURLExpiration string

	// server
	ServerExternalHostname string
	ServerHTTP2            string
	ServerHTTP2Bind        string
	ServerHTTP3            string
	ServerHTTP3Bind        string
	ServerMinifyHTML       string
	ServerRoles            string
	ServerTLSCertPath      string
	ServerTLSKeyPath       string

	// user
	UserEmail    string
	UserGroups   string
	UserPassword string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName: "application-name",
	SoftwareVersion: "software-version", // Set at build
	TokenSalt:       "token-salt",

	// database
	DbType:          "db-type",
	DbAddress:       "db-address",
	DbPort:          "db-port",
	DbUser:          "db-user",
	DbPassword:      "db-password",
	DbDatabase:      "db-database",
	DbTLSMode:       "db-tls-mode",
	DbTLSCACert:     "db-tls-ca-cert",
	DbLoadTestData:  "test-data", // CLI only
	DbEncryptionKey: "db-crypto-key",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// auth
	AccessExpiration:  "access-expiration",
	AccessSecret:      "access-secret",
	RefreshExpiration: "refresh-expiration",
	RefreshSecret:     "refresh-secret",

	// s3
	S3Endpoint:               "s3-endpoint",
	S3Region:                 "s3-region",
	S3AccessKeyID:            "s3-access-id",
	S3SecretAccessKey:        "s3-secret-key",
	S3UseSSL:                 "s3-ssl",
	S3Bucket:                 "s3-bucket",
	S3PresignedURLExpiration: "s3-url-expiration",

	// server
	ServerExternalHostname: "external-hostname",
	ServerHTTP2:            "http2",
	ServerHTTP2Bind:        "http2-bind",
	ServerHTTP3:            "http3",
	ServerHTTP3Bind:        "http3-bind",
	ServerMinifyHTML:       "minify-html",
	ServerRoles:            "server-role",
	ServerTLSCertPath:      "tls-cert",
	ServerTLSKeyPath:       "tls-key",

	// user
	UserEmail:    "email",
	UserGroups:   "group",
	UserPassword: "password",
}
