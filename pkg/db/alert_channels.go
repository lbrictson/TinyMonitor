package db

import (
	"context"
	"errors"
	"github.com/lbrictson/TinyMonitor/ent"
	"github.com/lbrictson/TinyMonitor/ent/alertchannel"
	"github.com/lbrictson/TinyMonitor/pkg/validators"
	"time"
)

type BaseAlertChannel struct {
	Name        string                 `json:"name"`
	ChannelType string                 `json:"channel_type"`
	Config      map[string]interface{} `json:"config"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	Monitors    []string               `json:"monitors"`
}

func convertEntAlertChannelToDBAlertChannel(entAlertChannel *ent.AlertChannel) *BaseAlertChannel {
	if entAlertChannel == nil {
		return nil
	}
	var monitors []string
	for _, monitor := range entAlertChannel.Edges.Monitors {
		monitors = append(monitors, monitor.ID)
	}
	return &BaseAlertChannel{
		Name:        entAlertChannel.ID,
		ChannelType: entAlertChannel.AlertChannelType,
		Config:      entAlertChannel.Config,
		CreatedAt:   entAlertChannel.CreatedAt,
		UpdatedAt:   entAlertChannel.UpdatedAt,
		Monitors:    monitors,
	}
}

type CreateAlertChannelInput struct {
	Name        string                 `json:"name"`
	ChannelType string                 `json:"channel_type"`
	Config      map[string]interface{} `json:"config"`
}

func (i *CreateAlertChannelInput) validate() error {
	if validators.ValidateName(i.Name) != nil {
		return errors.New("invalid name")
	}
	if validators.ValidateAlertChannelType(i.ChannelType) != nil {
		return errors.New("invalid alert channel type")
	}
	return nil
}

func (db *DatabaseConnection) CreateAlertChannel(ctx context.Context, input CreateAlertChannelInput) (*BaseAlertChannel, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	entAlertChannel, err := db.client.AlertChannel.Create().
		SetID(input.Name).
		SetAlertChannelType(input.ChannelType).
		SetConfig(input.Config).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntAlertChannelToDBAlertChannel(entAlertChannel), nil
}

type UpdateAlertChannelInput struct {
	Config map[string]interface{} `json:"config"`
}

func (i *UpdateAlertChannelInput) validate() error {
	return nil
}

func (db *DatabaseConnection) UpdateAlertChannel(ctx context.Context, name string, input UpdateAlertChannelInput) (*BaseAlertChannel, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	entAlertChannel, err := db.client.AlertChannel.UpdateOneID(name).
		SetConfig(input.Config).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntAlertChannelToDBAlertChannel(entAlertChannel), nil
}

func (db *DatabaseConnection) DeleteAlertChannel(ctx context.Context, name string) error {
	return db.client.AlertChannel.DeleteOneID(name).Exec(ctx)
}

func (db *DatabaseConnection) GetAlertChannel(ctx context.Context, name string) (*BaseAlertChannel, error) {
	entAlertChannel, err := db.client.AlertChannel.Query().WithMonitors().Where(alertchannel.ID(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntAlertChannelToDBAlertChannel(entAlertChannel), nil
}

func (db *DatabaseConnection) ListAlertChannels(ctx context.Context) ([]*BaseAlertChannel, error) {
	entAlertChannels, err := db.client.AlertChannel.Query().WithMonitors().All(ctx)
	if err != nil {
		return nil, err
	}
	var alertChannels []*BaseAlertChannel
	for _, entAlertChannel := range entAlertChannels {
		alertChannels = append(alertChannels, convertEntAlertChannelToDBAlertChannel(entAlertChannel))
	}
	return alertChannels, nil
}
