package config

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName string
	SoftwareVersion string

	// database
	DbType      string
	DbAddress   string
	DbPort      int
	DbUser      string
	DbPassword  string
	DbDatabase  string
	DbTLSMode   string
	DbTLSCACert string

	// server
	ServerRoles []string
}

// Defaults contains the default values
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName: "megabot",

	// database
	DbType:      "postgres",
	DbAddress:   "",
	DbPort:      5432,
	DbUser:      "",
	DbPassword:  "",
	DbDatabase:  "megabot",
	DbTLSMode:   "disable",
	DbTLSCACert: "",

	// server
	ServerRoles: []string{ServerRoleGraphql},
}
