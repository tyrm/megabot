package config

import "time"

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName string
	SoftwareVersion string
	TokenSalt       string

	// database
	DbType          string
	DbAddress       string
	DbPort          int
	DbUser          string
	DbPassword      string
	DbDatabase      string
	DbTLSMode       string
	DbTLSCACert     string
	DbLoadTestData  bool
	DbEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       int
	RedisPassword string

	// auth
	AccessExpiration  time.Duration
	AccessSecret      string
	RefreshExpiration time.Duration
	RefreshSecret     string

	// s3
	S3Endpoint               string
	S3Region                 string
	S3AccessKeyID            string
	S3SecretAccessKey        string
	S3UseSSL                 bool
	S3Bucket                 string
	S3PresignedURLExpiration time.Duration

	// server
	ServerExternalHostname string
	ServerHTTP2            bool
	ServerHTTP2Bind        string
	ServerHTTP3            bool
	ServerHTTP3Bind        string
	ServerMinifyHTML       bool
	ServerRoles            []string
	ServerTLSCertPath      string
	ServerTLSKeyPath       string

	// user
	UserEmail    string
	UserGroups   []string
	UserPassword string
}

// Defaults contains the default values
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName: "megabot",

	// database
	DbType:         "postgres",
	DbAddress:      "",
	DbPort:         5432,
	DbUser:         "",
	DbPassword:     "",
	DbDatabase:     "megabot",
	DbTLSMode:      "disable",
	DbTLSCACert:    "",
	DbLoadTestData: false,

	// redis
	RedisAddress:  "localhost:6379",
	RedisDB:       0,
	RedisPassword: "",

	// auth
	AccessExpiration:  time.Minute * 15,
	RefreshExpiration: time.Hour * 24 * 7,

	// s3
	S3Endpoint:               "play.min.io",
	S3Region:                 "us-east-1",
	S3UseSSL:                 true,
	S3Bucket:                 "megabot",
	S3PresignedURLExpiration: 10 * time.Second,

	// server
	ServerExternalHostname: "localhost",
	ServerHTTP2:            true,
	ServerHTTP2Bind:        ":5000",
	ServerHTTP3:            false,
	ServerHTTP3Bind:        ":5000",
	ServerMinifyHTML:       true,
	ServerRoles: []string{
		ServerRoleGraphQL,
		ServerRoleWebapp,
	},
	ServerTLSCertPath: "server.crt",
	ServerTLSKeyPath:  "server.key",
}
