package api

import "errors"

type InfluxDBV1SinkConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type TimeStreamSinkConfig struct {
	Region             string `json:"region"`
	AWSAccessKeyID     string `json:"aws_access_key_id"`
	AWSSecretAccessKey string `json:"aws_secret_access_key"`
	DBName             string `json:"db_name"`
	TableName          string `json:"table_name"`
}

type CloudWatchSinkConfig struct {
	Region             string `json:"region"`
	AWSAccessKeyID     string `json:"aws_access_key_id"`
	AWSSecretAccessKey string `json:"aws_secret_access_key"`
}

func validateInfluxDBV1SinkConfig(config InfluxDBV1SinkConfig) error {
	if config.Host == "" {
		return errors.New("host is required for InfluxDBV1 sink")
	}
	if config.Username == "" {
		return errors.New("username is required for InfluxDBV1 sink")
	}
	if config.Password == "" {
		return errors.New("password is required for InfluxDBV1 sink")
	}
	if config.Database == "" {
		return errors.New("database is required for InfluxDBV1 sink")
	}
	return nil
}

func validateTimeStreamSinkConfig(config TimeStreamSinkConfig) error {
	if config.DBName == "" {
		return errors.New("db_name is required for TimeStream sink")
	}
	if config.TableName == "" {
		return errors.New("table_name is required for TimeStream sink")
	}
	return nil
}
