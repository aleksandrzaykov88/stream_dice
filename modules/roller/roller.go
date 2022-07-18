package roller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func Throw(c echo.Context) error {
	return c.String(http.StatusOK, "You roll 20")
}
