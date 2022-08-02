package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Healthcheck func is simple endpoint for the echo app, primarily useful
// if the app is deployed on kubernetes
func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "I am OK, don't worry :)")
}
