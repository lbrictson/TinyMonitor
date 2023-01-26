package sink

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"strconv"
)

type TimeStreamSink struct {
	ts        *timestreamwrite.TimestreamWrite
	dbName    string
	tableName string
}

type NewTimeStreamSinkInput struct {
	DatabaseName string
	TableName    string
	Region       string
	AWSSecretKey string
	AWSAccessKey string
}

func NewTimeStreamSink(input NewTimeStreamSinkInput) (*TimeStreamSink, error) {
	if input.DatabaseName == "" {
		return nil, errors.New("database_name is required for TimeStream sink")
	}
	if input.TableName == "" {
		return nil, errors.New("table_name is required for TimeStream sink")
	}
	if input.Region == "" {
		input.Region = "us-east-1"
	}
	writeSvc := &timestreamwrite.TimestreamWrite{}
	if input.AWSSecretKey == "" {
		// Use AWS credentials from environment variables
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(input.Region),
		})
		if err != nil {
			return nil, err
		}
		writeSvc = timestreamwrite.New(sess)
	} else {
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(input.AWSAccessKey, input.AWSSecretKey, ""),
			Region:      &input.Region,
		})
		if err != nil {
			return nil, err
		}
		writeSvc = timestreamwrite.New(sess)
	}
	return &TimeStreamSink{
		ts:        writeSvc,
		dbName:    input.DatabaseName,
		tableName: input.TableName,
	}, nil
}

func (s *TimeStreamSink) SendMetric(input []SendMetricInput) error {
	records := []*timestreamwrite.Record{}
	for _, metric := range input {
		records = append(records, &timestreamwrite.Record{
			Dimensions: []*timestreamwrite.Dimension{
				{
					Name:  aws.String("TinyStatus"),
					Value: aws.String(metric.MonitorName),
				},
			},
			MeasureName:      aws.String(metric.MetricName),
			MeasureValue:     aws.String(strconv.FormatFloat(metric.MetricValue, 'f', -1, 64)),
			MeasureValueType: aws.String("DOUBLE"),
			Time:             aws.String(strconv.FormatInt(metric.MetricTime.UnixNano(), 10)),
			TimeUnit:         aws.String("NANOSECONDS"),
		})
	}
	writeRecordsInput := &timestreamwrite.WriteRecordsInput{
		DatabaseName: aws.String(s.dbName),
		TableName:    aws.String(s.tableName),
		Records:      records,
	}

	_, err := s.ts.WriteRecords(writeRecordsInput)
	if err != nil {
		return fmt.Errorf("error writing to TimeStream: %v", err)
	}
	return nil
}
