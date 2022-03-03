package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName string
	SoftwareVersion string

	// database
	DbType      string
	DbAddress   string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbDatabase  string
	DbTLSMode   string
	DbTLSCACert string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// auth
	AccessExpiration  string
	AccessSecret      string
	RefreshExpiration string
	RefreshSecret     string

	// database
	DatabaseTestData string

	// server
	ServerRoles string

	// user
	UserEmail    string
	UserGroups   string
	UserPassword string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path",
	LogLevel:   "log-level",

	// application
	ApplicationName: "application-name",
	SoftwareVersion: "software-version",

	// database
	DbType:      "db-type",
	DbAddress:   "db-address",
	DbPort:      "db-port",
	DbUser:      "db-user",
	DbPassword:  "db-password",
	DbDatabase:  "db-database",
	DbTLSMode:   "db-tls-mode",
	DbTLSCACert: "db-tls-ca-cert",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// auth
	AccessExpiration:  "access-expiration",
	AccessSecret:      "access-secret",
	RefreshExpiration: "refresh-expiration",
	RefreshSecret:     "refresh-secret",

	// database
	DatabaseTestData: "test-data",

	// server
	ServerRoles: "server-role",

	// user
	UserEmail:    "email",
	UserGroups:   "group",
	UserPassword: "password",
}
