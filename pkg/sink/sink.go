package sink

import "time"

type SendMetricInput struct {
	Tags        map[string]string
	MetricName  string
	MetricValue float64
	MetricUnit  string
	MetricTime  time.Time
	MonitorName string
}

type Sink interface {
	SendMetric(input []SendMetricInput) error
}
