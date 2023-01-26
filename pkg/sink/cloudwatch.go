package sink

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type NewCloudWatchSinkInput struct {
	Region       string
	AWSSecretKey string
	AWSAccessKey string
}

type CloudWatchSink struct {
	cw *cloudwatch.CloudWatch
}

func NewCloudWatchSink(input NewCloudWatchSinkInput) (*CloudWatchSink, error) {
	// Connect to AWS Cloudwatch
	cw := &cloudwatch.CloudWatch{}
	if input.Region == "" {
		input.Region = "us-east-1"
	}
	if input.AWSSecretKey == "" {
		// Use AWS credentials from environment variables
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(input.Region),
		})
		if err != nil {
			return nil, err
		}
		cw = cloudwatch.New(sess)
	} else {
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(input.AWSAccessKey, input.AWSSecretKey, ""),
			Region:      &input.Region,
		})
		if err != nil {
			return nil, err
		}
		cw = cloudwatch.New(sess)
	}
	return &CloudWatchSink{
		cw: cw,
	}, nil
}

func (s *CloudWatchSink) SendMetric(input []SendMetricInput) error {
	metrics := []*cloudwatch.MetricDatum{}
	for _, metric := range input {
		metrics = append(metrics, &cloudwatch.MetricDatum{
			MetricName: aws.String(fmt.Sprintf("%v|%v", metric.MonitorName, metric.MetricName)),
			Unit:       aws.String(metric.MetricUnit),
			Value:      aws.Float64(metric.MetricValue),
			Timestamp:  aws.Time(metric.MetricTime),
		})
	}
	// Send metrics to AWS Cloudwatch
	_, err := s.cw.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace:  aws.String("Custom/TinyStatus"),
		MetricData: metrics,
	})
	return err
}
