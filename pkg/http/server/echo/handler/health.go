package handler

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/http/server/echo/parser"
	"github.com/jpfaria/goignite/pkg/http/server/service"
	"github.com/labstack/echo/v4"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {
	return parser.JSONResponse(c, http.StatusOK, service.Health(c.Request().Context()), nil)
}
