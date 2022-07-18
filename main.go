package main

import (
	"fmt"

	"github.com/aleksandrzaykov88/stream_dice/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Start API
	api.GetEcho(e)
	api.ServeRoutes()

	// Start server
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}
