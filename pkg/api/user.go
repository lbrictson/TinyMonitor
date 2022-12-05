package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserModel struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	ID        int       `json:"id"`
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
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		LockedOut: user.Locked,
	}
}

func (s *Server) getUser(c echo.Context) error {
	param := c.Param("id")
	// Convert param to int
	id, err := strconv.Atoi(param)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, errors.New("invalid user id"))
	}
	// Get user from database
	user, err := s.DBConnection.GetUserByID(c.Request().Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("user not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBUserToAPIUser(user))
}
