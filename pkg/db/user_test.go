package db

import (
	"context"
	"testing"
)

func TestUserHappyPaths(t *testing.T) {
	ctx := context.TODO()
	dbConn, err := NewDatabaseConnection(NewDatabaseConnectionInput{
		InMemory: true,
		Location: "",
	})
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
		return
	}
	// Create a user
	user, err := dbConn.CreateUser(ctx, CreateUserInput{
		Username: "test",
		APIKey:   "testkey",
		Role:     "admin",
	})
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
		return
	}
	// List users
	users, err := dbConn.ListUsers(ctx)
	if err != nil {
		t.Fatalf("Error listing users: %v", err)
		return
	}
	if len(users) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(users))
	}
	// Get single user by username
	u, err := dbConn.GetUserByUsername(ctx, "test")
	if err != nil {
		t.Fatalf("Error getting user by username: %v", err)
		return
	}
	if u.Username != "test" {
		t.Fatalf("Expected username to be 'test', got '%s'", u.Username)
	}
	// Get single user by ID
	u, err = dbConn.GetUserByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("Error getting user by ID: %v", err)
		return
	}
	if u.Username != "test" {
		t.Fatalf("Expected username to be 'test', got '%s'", u.Username)
	}
	// Update user
	u.Locked = true
	u.Role = "read_only"
	u.APIKey = "newkey"
	u, err = dbConn.UpdateUser(ctx, u)
	if err != nil {
		t.Fatalf("Error updating user: %v", err)
		return
	}
	// Get updated user
	u, err = dbConn.GetUserByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("Error getting user by ID: %v", err)
		return
	}
	if u.Username != "test" {
		t.Fatalf("Expected username to be 'test', got '%s'", u.Username)
	}
	if u.Role != "read_only" {
		t.Fatalf("Expected role to be 'read_only', got '%s'", u.Role)
	}
	if u.APIKey != "newkey" {
		t.Fatalf("Expected API key to be 'newkey'")
	}
	if !u.Locked {
		t.Fatalf("Expected user to be locked")
	}
	// Delete user
	if err := dbConn.DeleteUser(ctx, user.ID); err != nil {
		t.Fatalf("Error deleting user: %v", err)
		return
	}
}
