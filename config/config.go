package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost    string
	UserServicePort    int
	ProductServiceHost string
	ProductServicePort int

	// context timeout in seconds
	CtxTimeout int

	// jwt ...
	SigningKey         string
	AccessTokenTimeout int

	// casbin...
	AuthConfigPath string
	CSVFilePath    string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	// another services
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 1111))
	c.ProductServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "127.0.0.1"))
	c.ProductServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 1112))

	// casbin...
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_PATH", "./config/auth.conf"))
	c.CSVFilePath = cast.ToString(getOrReturnDefault("CSV_FILE_PATH", "./config/casbin_rules.csv"))

	// jwt
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.SigningKey = cast.ToString(getOrReturnDefault("SIGNING_KEY", "test-key"))
	c.AccessTokenTimeout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMOUT", 3600))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
