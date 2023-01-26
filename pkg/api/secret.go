package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/validators"
	"net/http"
	"time"
)

type SecretModel struct {
	Name          string    `json:"name"`
	LastUpdatedBy string    `json:"last_updated_by"`
	UpdatedAt     time.Time `json:"last_updated"`
	CreatedAt     time.Time `json:"created_at"`
}

func convertDBSecretToAPISecret(dbSecret *db.Secret) *SecretModel {
	if dbSecret == nil {
		return nil
	}
	return &SecretModel{
		Name:          dbSecret.Name,
		LastUpdatedBy: dbSecret.CreatedBy,
		UpdatedAt:     dbSecret.UpdatedAt,
		CreatedAt:     dbSecret.CreatedAt,
	}
}

func (s *Server) getSecret(c echo.Context) error {
	// Get the secret from the database
	secret, err := s.dbConnection.GetSecretByName(c.Request().Context(), c.Param("id"))
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBSecretToAPISecret(secret))
}

type CreateSecretInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (data *CreateSecretInput) validate() error {
	if validators.ValidateName(data.Name) != nil {
		return validators.ValidateName(data.Name)
	}
	if data.Name == "" {
		return errors.New("name is required")
	}
	if data.Value == "" {
		return errors.New("value is required")
	}
	return nil
}

func (s *Server) createSecret(c echo.Context) error {
	data := CreateSecretInput{}
	if err := c.Bind(&data); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := data.validate(); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	// Make sure the secret doesn't already exist
	secret, _ := s.dbConnection.GetSecretByName(c.Request().Context(), data.Name)
	if secret != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, errors.New("secret already exists"))
	}
	_, err := s.dbConnection.CreateSecret(c.Request().Context(), db.CreateSecretInput{
		Name:      data.Name,
		Value:     data.Value,
		CreatedBy: c.Get("username").(string),
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusCreated, nil)
}

type UpdateSecretInput struct {
	Value string `json:"value"`
}

func (data *UpdateSecretInput) validate() error {
	if data.Value == "" {
		return errors.New("value is required")
	}
	return nil
}

func (s *Server) updateSecret(c echo.Context) error {
	data := UpdateSecretInput{}
	if err := c.Bind(&data); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := data.validate(); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	_, err := s.dbConnection.UpdateSecret(c.Request().Context(), db.UpdateSecretInput{
		Name:      c.Param("id"),
		Value:     data.Value,
		CreatedBy: c.Get("username").(string),
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, nil)
}

func (s *Server) deleteSecret(c echo.Context) error {
	err := s.dbConnection.DeleteSecret(c.Request().Context(), c.Param("id"))
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, nil)
}

func (s *Server) listSecrets(c echo.Context) error {
	secrets, err := s.dbConnection.ListSecrets(c.Request().Context())
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	data := []SecretModel{}
	for _, secret := range secrets {
		data = append(data, *convertDBSecretToAPISecret(secret))
	}
	return s.returnSuccessResponse(c, http.StatusOK, data)
}
