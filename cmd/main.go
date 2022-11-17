package main

import (
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"good-to-go/configs"
	"good-to-go/pkg/http/rest/handler"
	"good-to-go/pkg/util/auth"
)

func main() {
	configPaths := []string{"./configs"}
	err := configs.LoadConfig(configPaths)
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	c := jaegertracing.New(e, nil)
	defer c.Close()

	defaultLogger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, \n",
	})
	e.Use(defaultLogger)
	e.Use(middleware.Recover())

	e.GET("/healthcheck", handler.Healthcheck)

	api := e.Group("/api/v1")
	{
		config := middleware.JWTConfig{
			KeyFunc: auth.GetKey,
		}
		api.Use(middleware.JWTWithConfig(config))
		api.PUT("/upload", handler.DummyApi)
	}

	e.Logger.Fatal(e.Start(":" + configs.Configuration.Port))
}
