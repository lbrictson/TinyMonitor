package config

import "github.com/kelseyhightower/envconfig"

type ServerEnvConfig struct {
	Port       int    `envconfig:"PORT" default:"8080"`
	Hostname   string `envconfig:"DOMAIN" default:"localhost"`
	AutoTLS    bool   `envconfig:"AUTO_TLS" default:"false"`
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	LogFormat  string `envconfig:"LOG_FORMAT" default:"text"`
	DBLocation string `envconfig:"DB_LOCATION" default:"tmp/"`
}

func ReadServerConfig() (ServerEnvConfig, error) {
	s := &ServerEnvConfig{}
	err := envconfig.Process("TINYMONITOR", s)
	return *s, err
}
