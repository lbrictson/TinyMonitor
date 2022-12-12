package db

import (
	"context"
	"errors"
	"github.com/lbrictson/TinyMonitor/ent"
	"github.com/lbrictson/TinyMonitor/ent/user"
	"strings"
	"time"
)

type User struct {
	Username    string     `json:"username"`
	APIKey      string     `json:"api_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Role        string     `json:"role"`
	Locked      bool       `json:"locked"`
	LockedUntil *time.Time `json:"locked_until"`
}

func convertEntUserToDBUser(entUser *ent.User) *User {
	if entUser == nil {
		return nil
	}
	return &User{
		Username:    entUser.ID,
		APIKey:      entUser.APIKey,
		CreatedAt:   entUser.CreatedAt,
		UpdatedAt:   entUser.UpdatedAt,
		Role:        entUser.Role,
		Locked:      entUser.Locked,
		LockedUntil: entUser.LockedUntil,
	}
}

func (db *DatabaseConnection) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	u, err := db.client.User.Query().Where(user.ID(username)).First(ctx)
	return convertEntUserToDBUser(u), err
}

func (db *DatabaseConnection) ListUsers(ctx context.Context) ([]*User, error) {
	users, err := db.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var dbUsers []*User
	for _, u := range users {
		dbUsers = append(dbUsers, convertEntUserToDBUser(u))
	}
	return dbUsers, nil
}

func (db *DatabaseConnection) DeleteUser(ctx context.Context, username string) error {
	return db.client.User.DeleteOneID(username).Exec(ctx)
}

type CreateUserInput struct {
	Username string
	APIKey   string
	Role     string
}

func (i CreateUserInput) validate() error {
	if i.Username == "" {
		return errors.New("username is required")
	}
	if strings.Contains(i.Username, " ") {
		return errors.New("username cannot contain spaces")
	}
	// Validate username only contains letters, numbers, and dashes
	if !strings.ContainsAny(i.Username, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-") {
		return errors.New("username can only contain letters, numbers, and dashes")
	}
	if i.APIKey == "" {
		return errors.New("api_key is required")
	}
	if i.Role == "" {
		return errors.New("role is required")
	}
	return validateRole(i.Role)
}

func validateRole(role string) error {
	acceptableRoles := []string{"read_only", "write", "admin"}
	for _, r := range acceptableRoles {
		if role == r {
			return nil
		}
	}
	return errors.New("role must be one of: read_only, write, admin")
}

func (db *DatabaseConnection) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	u, err := db.client.User.Create().SetID(input.Username).SetAPIKey(input.APIKey).SetRole(input.Role).Save(ctx)
	return convertEntUserToDBUser(u), err
}

func (db *DatabaseConnection) UpdateUser(ctx context.Context, user *User) (*User, error) {
	if user == nil {
		return nil, errors.New("cannot update nil user")
	}
	if err := validateRole(user.Role); err != nil {
		return nil, err
	}
	q := db.client.User.UpdateOneID(user.Username).SetAPIKey(user.APIKey).SetRole(user.Role).SetLocked(user.Locked)
	if user.LockedUntil != nil {
		q.SetLockedUntil(*user.LockedUntil)
	}
	u, err := q.Save(ctx)
	return convertEntUserToDBUser(u), err
}
