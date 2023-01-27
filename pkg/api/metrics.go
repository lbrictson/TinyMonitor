package api

import (
	"context"
	"encoding/json"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/sink"
	"os"
	"strings"
	"sync"
)

var metricLock sync.Mutex
var metrics = make(map[string]sink.Sink)

func (s *Server) loadMetrics() error {
	m, err := s.dbConnection.ListSinks(context.Background())
	if err != nil {
		return err
	}
	for _, r := range m {
		err = loadSingleSinkIntoMetrics(*convertDBSinkToAPISink(r), s.dbConnection)
		if err != nil {
			return err
		}
	}
	return nil
}

func removeSinkFromMetrics(name string) {
	metricLock.Lock()
	defer metricLock.Unlock()
	delete(metrics, name)
}

func loadSingleSinkIntoMetrics(s BaseSink, databaseConn *db.DatabaseConnection) error {
	metricLock.Lock()
	defer metricLock.Unlock()
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	b, err := json.Marshal(s.Config)
	if err != nil {
		return err
	}
	switch strings.ToLower(s.SinkType) {
	case "influxdb-v1":
		conf := InfluxDBV1SinkConfig{}
		err = json.Unmarshal(b, &conf)
		if err != nil {
			return err
		}
		m, err := sink.NewInfluxDBV1Sink(sink.NewInfluxDBV1SinkInput{
			DatabaseName:   injectSecretsIntoContent(context.TODO(), databaseConn, conf.Database),
			Username:       injectSecretsIntoContent(context.TODO(), databaseConn, conf.Username),
			Password:       injectSecretsIntoContent(context.TODO(), databaseConn, conf.Password),
			ServerURL:      injectSecretsIntoContent(context.TODO(), databaseConn, conf.Host),
			SenderHostname: hostname,
		})
		if err != nil {
			return err
		}
		metrics[s.Name] = m
	case "cloudwatch":
		conf := CloudWatchSinkConfig{}
		err = json.Unmarshal(b, &conf)
		if err != nil {
			return err
		}
		m, err := sink.NewCloudWatchSink(sink.NewCloudWatchSinkInput{
			Region:       injectSecretsIntoContent(context.TODO(), databaseConn, conf.Region),
			AWSSecretKey: injectSecretsIntoContent(context.TODO(), databaseConn, conf.AWSSecretAccessKey),
			AWSAccessKey: injectSecretsIntoContent(context.TODO(), databaseConn, conf.AWSAccessKeyID),
		})
		if err != nil {
			return err
		}
		metrics[s.Name] = m
	case "timestream":
		conf := TimeStreamSinkConfig{}
		err = json.Unmarshal(b, &conf)
		if err != nil {
			return err
		}
		m, err := sink.NewTimeStreamSink(sink.NewTimeStreamSinkInput{
			DatabaseName: injectSecretsIntoContent(context.TODO(), databaseConn, conf.DBName),
			TableName:    injectSecretsIntoContent(context.TODO(), databaseConn, conf.TableName),
			Region:       injectSecretsIntoContent(context.TODO(), databaseConn, conf.Region),
			AWSSecretKey: injectSecretsIntoContent(context.TODO(), databaseConn, conf.AWSSecretAccessKey),
			AWSAccessKey: injectSecretsIntoContent(context.TODO(), databaseConn, conf.AWSAccessKeyID),
		})
		if err != nil {
			return err
		}
		metrics[s.Name] = m
	}
	return nil
}