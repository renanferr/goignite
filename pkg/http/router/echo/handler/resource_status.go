package handler

import (
	"net/http"

	"github.com/b2wdigital/goignite/pkg/http/router/echo/parser"
	"github.com/b2wdigital/goignite/pkg/http/router/service"
	"github.com/labstack/echo/v4"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c echo.Context) error {
	return parser.JSONResponse(c, http.StatusOK, service.ResourceStatus(c.Request().Context()), nil)
}
