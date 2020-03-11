package echo

import (
	"net/http"

	"github.com/b2wdigital/goignite/pkg/server/http/router"
	"github.com/labstack/echo/v4"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c echo.Context) error {
	return echo2.JSONResponse(c, http.StatusOK, router.ResourceStatus(c.Request().Context()), nil)
}
