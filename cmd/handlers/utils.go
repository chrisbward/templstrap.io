package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) (err error) {
	return component.Render(c.Request().Context(), c.Response())
}
