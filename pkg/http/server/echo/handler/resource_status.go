package handler

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/http/server/echo/parser"
	"github.com/jpfaria/goignite/pkg/http/server/service"
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
