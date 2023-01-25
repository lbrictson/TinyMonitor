package db

import (
	"context"
	"github.com/lbrictson/TinyMonitor/ent"
	"time"
)

type Secret struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
}

func convertEntSecretToDBSecret(entSecret *ent.Secret) *Secret {
	if entSecret == nil {
		return nil
	}
	return &Secret{
		Name:      entSecret.ID,
		Value:     entSecret.Value,
		CreatedAt: entSecret.CreatedAt,
		UpdatedAt: entSecret.UpdatedAt,
		CreatedBy: entSecret.CreatedBy,
	}
}

func (db *DatabaseConnection) GetSecretByName(ctx context.Context, name string) (*Secret, error) {
	s, err := db.client.Secret.Get(ctx, name)
	return convertEntSecretToDBSecret(s), err
}

type CreateSecretInput struct {
	Name      string
	Value     string
	CreatedBy string
}

func (db *DatabaseConnection) CreateSecret(ctx context.Context, input CreateSecretInput) (*Secret, error) {
	s, err := db.client.Secret.Create().SetID(input.Name).SetValue(input.Value).SetCreatedBy(input.CreatedBy).Save(ctx)
	return convertEntSecretToDBSecret(s), err
}

type UpdateSecretInput struct {
	Name      string
	Value     string
	CreatedBy string
}

func (db *DatabaseConnection) UpdateSecret(ctx context.Context, input UpdateSecretInput) (*Secret, error) {
	s, err := db.client.Secret.UpdateOneID(input.Name).SetValue(input.Value).SetCreatedBy(input.CreatedBy).Save(ctx)
	return convertEntSecretToDBSecret(s), err
}

func (db *DatabaseConnection) DeleteSecret(ctx context.Context, name string) error {
	return db.client.Secret.DeleteOneID(name).Exec(ctx)
}

func (db *DatabaseConnection) ListSecrets(ctx context.Context) ([]*Secret, error) {
	secrets, err := db.client.Secret.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var dbSecrets []*Secret
	for _, s := range secrets {
		dbSecrets = append(dbSecrets, convertEntSecretToDBSecret(s))
	}
	return dbSecrets, nil
}
