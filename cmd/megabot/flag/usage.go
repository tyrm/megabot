package flag

import "github.com/tyrm/megabot/internal/config"

var usage = config.KeyNames{
	ConfigPath: "Path to a file containing megabot configuration. Values set in this file will be overwritten by values set as env vars or arguments",
	LogLevel:   "Log level to run at: [trace, debug, info, warn, fatal]",

	// application
	ApplicationName: "Name of the application, used in various places internally",

	// database
	DbType:      "Database type: eg., postgres",
	DbAddress:   "Database ipv4 address, hostname, or filename",
	DbPort:      "Database port",
	DbUser:      "Database username",
	DbPassword:  "Database password",
	DbDatabase:  "Database name",
	DbTLSMode:   "Database tls mode",
	DbTLSCACert: "Path to CA cert for db tls connection",

	// server
	ServerRoles: "Server roles that should be started: [graphql]",
}
