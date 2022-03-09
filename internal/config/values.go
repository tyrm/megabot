package config

import "time"

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName string
	SoftwareVersion string

	// database
	DbType         string
	DbAddress      string
	DbPort         int
	DbUser         string
	DbPassword     string
	DbDatabase     string
	DbTLSMode      string
	DbTLSCACert    string
	DbLoadTestData bool

	// redis
	RedisAddress  string
	RedisDB       int
	RedisPassword string

	// auth
	AccessExpiration  time.Duration
	AccessSecret      string
	RefreshExpiration time.Duration
	RefreshSecret     string

	// server
	ServerExternalHostname string
	ServerRoles            []string

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

	// server
	ServerExternalHostname: "localhost",
	ServerRoles: []string{
		ServerRoleGraphQL,
		ServerRoleWebapp,
	},
}
