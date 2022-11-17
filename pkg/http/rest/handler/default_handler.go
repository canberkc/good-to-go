package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "I am OK, don't worry :)")
}

func DummyApi(c echo.Context) error {
	return c.String(http.StatusOK, "Yes your api call is OK")
}
