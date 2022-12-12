package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/security"
	"net/http"
	"strings"
	"time"
)

type UserModel struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LockedOut bool      `json:"locked_out"`
}

func convertDBUserToAPIUser(user *db.User) *UserModel {
	if user == nil {
		return nil
	}
	return &UserModel{
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		LockedOut: user.Locked,
	}
}

func (s *Server) getUser(c echo.Context) error {
	// Get user from database
	user, err := s.dbConnection.GetUserByUsername(c.Request().Context(), c.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("user not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBUserToAPIUser(user))
}

func (s *Server) listUsers(c echo.Context) error {
	users, err := s.dbConnection.ListUsers(c.Request().Context())
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	var apiUsers []*UserModel
	for _, u := range users {
		apiUsers = append(apiUsers, convertDBUserToAPIUser(u))
	}
	return s.returnSuccessResponse(c, http.StatusOK, apiUsers)
}

func (s *Server) deleteUser(c echo.Context) error {
	// Get user from database
	user, err := s.dbConnection.GetUserByUsername(c.Request().Context(), c.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("user not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	// If the user requesting the action is the same as the one being deleted, return an error because that would be
	// a silly use case
	if user.Username == c.Get("username").(string) {
		return s.returnErrorResponse(c, http.StatusBadRequest, errors.New("cannot delete self"))
	}
	// Delete user
	err = s.dbConnection.DeleteUser(c.Request().Context(), user.Username)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, nil)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type CreateUserResponse struct {
	APIKey string `json:"api_key"`
}

func (s *Server) createUser(c echo.Context) error {
	// Get the user from the request
	var user CreateUserRequest
	err := c.Bind(&user)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	// Validate the user doesn't already exist
	_, err = s.dbConnection.GetUserByUsername(c.Request().Context(), user.Username)
	if err == nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, errors.New("user already exists"))
	}
	// Create the user in the database
	newUser, err := s.dbConnection.CreateUser(c.Request().Context(), db.CreateUserInput{
		Username: user.Username,
		APIKey:   security.GenerateAPIKey(),
		Role:     user.Role,
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, CreateUserResponse{APIKey: newUser.APIKey})
}

type UpdateUserRequest struct {
	Role      *string `json:"role,omitempty"`
	LockedOut *bool   `json:"locked_out,omitempty"`
}

func (s *Server) updateUser(c echo.Context) error {
	// Get the user from the request
	var user UpdateUserRequest
	err := c.Bind(&user)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	// Get user from database
	dbUser, err := s.dbConnection.GetUserByUsername(c.Request().Context(), c.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("user not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	// Update the user
	if user.Role != nil {
		dbUser.Role = *user.Role
	}
	if user.LockedOut != nil {
		dbUser.Locked = *user.LockedOut
	}
	updatedUser, err := s.dbConnection.UpdateUser(c.Request().Context(), dbUser)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBUserToAPIUser(updatedUser))
}
