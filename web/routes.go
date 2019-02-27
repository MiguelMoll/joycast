package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo) {
	e.GET("/", hello)
}

// Handler
func hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
