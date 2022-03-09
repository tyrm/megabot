package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName string
	SoftwareVersion string

	// database
	DbType         string
	DbAddress      string
	DbPort         string
	DbUser         string
	DbPassword     string
	DbDatabase     string
	DbTLSMode      string
	DbTLSCACert    string
	DbLoadTestData string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// auth
	AccessExpiration  string
	AccessSecret      string
	RefreshExpiration string
	RefreshSecret     string

	// server
	ServerExternalHostname string
	ServerRoles            string

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

	// database
	DbType:         "db-type",
	DbAddress:      "db-address",
	DbPort:         "db-port",
	DbUser:         "db-user",
	DbPassword:     "db-password",
	DbDatabase:     "db-database",
	DbTLSMode:      "db-tls-mode",
	DbTLSCACert:    "db-tls-ca-cert",
	DbLoadTestData: "test-data", // CLI only

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// auth
	AccessExpiration:  "access-expiration",
	AccessSecret:      "access-secret",
	RefreshExpiration: "refresh-expiration",
	RefreshSecret:     "refresh-secret",

	// server
	ServerExternalHostname: "external-hostname",
	ServerRoles:            "server-role",

	// user
	UserEmail:    "email",
	UserGroups:   "group",
	UserPassword: "password",
}
