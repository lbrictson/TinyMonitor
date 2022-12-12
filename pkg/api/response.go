package api

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Error string `json:"error"`
}

func (s *Server) returnErrorResponse(c echo.Context, errorCode int, err error) error {
	user := "anonymous"
	if u, ok := c.Get("username").(string); ok {
		user = u
	}
	s.logger.WithFields(map[string]interface{}{
		"error":       err,
		"status_code": errorCode,
		"path":        c.Path(),
		"method":      c.Request().Method,
		"user":        user,
	}).Error("returning error response")
	return c.JSON(errorCode, &ErrorResponse{
		Error: err.Error(),
	})
}

func (s *Server) returnSuccessResponse(c echo.Context, statusCode int, data interface{}) error {
	user := "anonymous"
	if u, ok := c.Get("username").(string); ok {
		user = u
	}
	s.logger.WithFields(map[string]interface{}{
		"status_code": statusCode,
		"path":        c.Path(),
		"method":      c.Request().Method,
		"user":        user,
	}).Info("returning success response")
	return c.JSON(statusCode, data)
}
