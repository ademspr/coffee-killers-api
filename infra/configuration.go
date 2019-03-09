package infra

import "os"

// Configuration Environment configurations
type Configuration struct {
	Environment  string
	DbConnection string
}

// GetConfigurations Get environment configurations.
func GetConfigurations() Configuration {
	dbConnection := os.Getenv("DBCONNECTIONSTR")
	env := os.Getenv("ENV")
	s := Configuration{env, dbConnection}
	return s
}
