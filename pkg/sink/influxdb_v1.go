package sink

import (
	influxClient "github.com/influxdata/influxdb1-client/v2"
)

type InfluxDBV1Sink struct {
	dbName   string
	c        influxClient.Client
	hostname string
}

type NewInfluxDBV1SinkInput struct {
	DatabaseName   string
	Username       string
	Password       string
	ServerURL      string
	SenderHostname string
}

func NewInfluxDBV1Sink(input NewInfluxDBV1SinkInput) (*InfluxDBV1Sink, error) {
	c, err := influxClient.NewHTTPClient(influxClient.HTTPConfig{
		Addr:     input.ServerURL,
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &InfluxDBV1Sink{
		dbName:   input.DatabaseName,
		c:        c,
		hostname: input.SenderHostname,
	}, nil
}

func (s *InfluxDBV1Sink) SendMetric(input []SendMetricInput) error {
	bp, err := influxClient.NewBatchPoints(influxClient.BatchPointsConfig{
		Database:  s.dbName,
		Precision: "ns",
	})
	if err != nil {
		return err
	}
	for _, metric := range input {
		tags := metric.Tags
		tags["host"] = s.hostname
		tags["monitor"] = metric.MonitorName
		pt, err := influxClient.NewPoint(metric.MetricName, tags, map[string]interface{}{
			metric.MetricName: metric.MetricValue,
		}, metric.MetricTime)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}
	return s.c.Write(bp)
}
