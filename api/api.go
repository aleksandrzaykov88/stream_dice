package api

import (
	"github.com/aleksandrzaykov88/stream_dice/modules/roller"
	"github.com/labstack/echo/v4"
)

var api *echo.Echo

func GetEcho(e *echo.Echo) {
	api = e
}

// Routes
func ServeRoutes() {
	api.GET("/roll", roller.Throw)
}
