package default_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "I am OK, don't worry :)")
}
