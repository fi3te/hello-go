package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setUpMiddleware(e *echo.Echo) {
	basicAuthConfig := middleware.BasicAuthConfig{
		Skipper:   skipAuthentication,
		Validator: authenticate,
	}
	e.Use(middleware.BasicAuthWithConfig(basicAuthConfig))
	e.Use(middleware.Recover())
}

func skipAuthentication(c echo.Context) bool {
	return !strings.Contains(c.Path(), "secret")
}

func authenticate(username string, password string, c echo.Context) (bool, error) {
	return username == "admin" && password == "password", nil
}
