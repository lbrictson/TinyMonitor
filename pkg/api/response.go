package api

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Error string `json:"error"`
}

func (s *Server) returnErrorResponse(c echo.Context, errorCode int, err error) error {
	s.logger.WithFields(map[string]interface{}{
		"error":       err,
		"status_code": errorCode,
		"path":        c.Path(),
		"method":      c.Request().Method,
	}).Error("returning error response")
	return c.JSON(errorCode, &ErrorResponse{
		Error: err.Error(),
	})
}

type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func (s *Server) returnSuccessResponse(c echo.Context, statusCode int, data interface{}) error {
	s.logger.WithFields(map[string]interface{}{
		"status_code": statusCode,
		"path":        c.Path(),
		"method":      c.Request().Method,
	}).Info("returning success response")
	return c.JSON(statusCode, data)
}
