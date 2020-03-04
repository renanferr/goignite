package handler

import (
	"github.com/b2wdigital/goignite/pkg/http/router/echo/parser"
	"github.com/b2wdigital/goignite/pkg/http/router/service"
	"github.com/labstack/echo/v4"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {

	resp, httpCode := service.Health(c.Request().Context())

	return parser.JSONResponse(c, httpCode, resp, nil)
}
