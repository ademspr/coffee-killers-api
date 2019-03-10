package infra

import (
	"os"
	"strconv"
)

// MongoConfiguration mongodb configuration
type MongoConfiguration struct {
	Host   string `json:"host"`
	DbName string `json:"dbName"`
}

// ServerConfiguration server configuration
type ServerConfiguration struct {
	Port int64 `json:"port"`
}

// ApplicationConfiguration application environment configuration
type ApplicationConfiguration struct {
	Environment string `json:"env"`
}

// Auth authorization configuration
type AuthConfiguration struct {
	Secret string `json:"secret"`
}

// Configuration Environment configurations
type Configuration struct {
	Mongo  *MongoConfiguration       `json:"mongo"`
	Server *ServerConfiguration      `json:"server"`
	Env    *ApplicationConfiguration `json:"app"`
	Auth   *AuthConfiguration        `json:"auth"`
}

// GetConfigurations get api configurations.
func GetConfigurations() *Configuration {

	port, _ := strconv.ParseInt(getConfigOrDefault("PORT", "8000"), 0, 64)

	return &Configuration{
		Mongo: &MongoConfiguration{
			Host:   getConfigOrDefault("DBHOST", "127.0.0.1:27017"),
			DbName: getConfigOrDefault("DBNAME", "coffee-killers")},
		Server: &ServerConfiguration{
			Port: port},
		Auth: &AuthConfiguration{
			Secret: getConfigOrDefault("SECRET", "devscret"),
		},
		Env: &ApplicationConfiguration{
			Environment: getConfigOrDefault("APP_ENV", "development")}}
}

func getConfigOrDefault(ev string, dv string) string {
	c := os.Getenv(ev)
	if c == "" {
		return dv
	}
	return c
}
