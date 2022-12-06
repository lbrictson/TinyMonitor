package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (s *Server) userAuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// validate the auth headers are present
		username := c.Request().Header.Get("X-Username")
		apiKey := c.Request().Header.Get("X-Api-Key")
		if username == "" || apiKey == "" {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		// validate the auth headers are valid
		// Check the cache first
		cachedAPIKey, ok := s.memoryCache.Get(username)
		if ok {
			// If the cached API key matches, we're good
			if cachedAPIKey == apiKey {
				c.Set("username", username)
				return next(c)
			} else {
				// The cached API key is invalid, so remove it from the cache
				s.memoryCache.Delete(username)
				return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
			}
		}
		// The username was not in the cache, we will need to hit the database file to validate the API key
		user, err := s.dbConnection.GetUserByUsername(c.Request().Context(), username)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusInternalServerError, errors.New("unauthorized"))
		}
		if user == nil {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		if user.APIKey != apiKey {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		// The API key is valid, so cache it
		s.memoryCache.Set(username, apiKey, 1*time.Hour)
		c.Set("username", username)
		return next(c)
	}
}

func (s *Server) adminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Get("username").(string)
		// Check the role cache
		cachedRole, ok := s.memoryCache.Get(username + "-role")
		if ok {
			// Cache hit
			if cachedRole == "admin" {
				c.Set("role", "admin")
				return next(c)
			} else {
				// The cached role is not admin, so remove it from the cache, this way on a future hit there's a chance
				// the user's role has been updated and it can be gotten from the DB if needed
				s.memoryCache.Delete(username + "-role")
				return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
			}
		}
		user, err := s.dbConnection.GetUserByUsername(c.Request().Context(), username)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusInternalServerError, errors.New("unauthorized"))
		}
		if user == nil {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		if user.Role != "admin" {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		// The user is an admin, so cache the role
		c.Set("role", "admin")
		s.memoryCache.Set(username+"-role", user.Role, 1*time.Hour)
		return next(c)
	}
}

func (s *Server) writeRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Get("username").(string)
		// Check the role cache
		cachedRole, ok := s.memoryCache.Get(username + "-role")
		if ok {
			// Cache hit
			if cachedRole == "admin" || cachedRole == "write" {
				c.Set("role", cachedRole)
				return next(c)
			} else {
				// The cached role is not admin or write, so remove it from the cache, this way on a future hit there's a chance
				// the user's role has been updated and it can be gotten from the DB if needed
				s.memoryCache.Delete(username + "-role")
				return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
			}
		}
		user, err := s.dbConnection.GetUserByUsername(c.Request().Context(), username)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusInternalServerError, errors.New("unauthorized"))
		}
		if user == nil {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		if user.Role != "admin" && user.Role != "write" {
			return s.returnErrorResponse(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		// The user is an admin or write, so cache the role
		s.memoryCache.Set(username+"-role", user.Role, 1*time.Hour)
		return next(c)
	}
}
