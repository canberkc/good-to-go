package main

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"good-to-go/pkg/http/rest/handler"
)

func main() {
	e := echo.New()
	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", handler.Healthcheck)

	e.Logger.Fatal(e.Start(":8080"))
}
