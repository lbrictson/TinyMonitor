package config

import "github.com/kelseyhightower/envconfig"

type ServerEnvConfig struct {
	Port      int    `envconfig:"PORT" default:"8080"`
	Hostname  string `envconfig:"DOMAIN" default:"localhost"`
	AutoTLS   bool   `envconfig:"AUTO_TLS" default:"false"`
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
	LogFormat string `envconfig:"LOG_FORMAT" default:"text"`
	DBPort    int    `envconfig:"DB_PORT" default:"5432"`
	DBHost    string `envconfig:"DB_HOST" default:"localhost"`
	DBUser    string `envconfig:"DB_USER" default:"postgres"`
	DBPass    string `envconfig:"DB_PASS" default:"postgres"`
	DBName    string `envconfig:"DB_NAME" default:"postgres"`
	DBSSLMode string `envconfig:"DB_SSL_MODE" default:"disable"`
}

func ReadServerConfig() (ServerEnvConfig, error) {
	s := &ServerEnvConfig{}
	err := envconfig.Process("TINYMONITOR", s)
	return *s, err
}
