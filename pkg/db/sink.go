package db

import (
	"context"
	"errors"
	"github.com/lbrictson/TinyMonitor/ent"
	"github.com/lbrictson/TinyMonitor/pkg/validators"
	"time"
)

type BaseSink struct {
	Name      string                 `json:"name"`
	SinkType  string                 `json:"sink_type"`
	Config    map[string]interface{} `json:"config"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

func convertEntSinkToDBSink(entSink *ent.Sink) *BaseSink {
	if entSink == nil {
		return nil
	}
	return &BaseSink{
		Name:      entSink.ID,
		SinkType:  entSink.SinkType,
		Config:    entSink.Config,
		CreatedAt: entSink.CreatedAt,
		UpdatedAt: entSink.UpdatedAt,
	}
}

func (db *DatabaseConnection) ListSinks(ctx context.Context) ([]*BaseSink, error) {
	entSinks, err := db.client.Sink.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var sinks []*BaseSink
	for _, entSink := range entSinks {
		sinks = append(sinks, convertEntSinkToDBSink(entSink))
	}
	return sinks, nil
}

type CreateSinkInput struct {
	Name     string                 `json:"name"`
	SinkType string                 `json:"sink_type"`
	Config   map[string]interface{} `json:"config"`
}

func (i *CreateSinkInput) validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if validators.ValidateSinkType(i.SinkType) != nil {
		return errors.New("invalid sink type")
	}
	return nil
}

func (db *DatabaseConnection) CreateSink(ctx context.Context, input CreateSinkInput) (*BaseSink, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	entSink, err := db.client.Sink.Create().
		SetID(input.Name).
		SetSinkType(input.SinkType).
		SetConfig(input.Config).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntSinkToDBSink(entSink), nil
}

type UpdateSinkInput struct {
	Config map[string]interface{} `json:"config"`
}

func (db *DatabaseConnection) UpdateSink(ctx context.Context, name string, input UpdateSinkInput) (*BaseSink, error) {
	entSink, err := db.client.Sink.UpdateOneID(name).
		SetConfig(input.Config).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntSinkToDBSink(entSink), nil
}

func (db *DatabaseConnection) DeleteSink(ctx context.Context, name string) error {
	err := db.client.Sink.DeleteOneID(name).Exec(ctx)
	return err
}

func (db *DatabaseConnection) GetSink(ctx context.Context, name string) (*BaseSink, error) {
	entSink, err := db.client.Sink.Get(ctx, name)
	if err != nil {
		return nil, err
	}
	return convertEntSinkToDBSink(entSink), nil
}
