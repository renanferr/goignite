package echo

import (
	"net/http"

	"github.com/b2wdigital/goignite/pkg/transport/server/http/router"
	"github.com/labstack/echo/v4"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c echo.Context) error {
	return JSONResponse(c, http.StatusOK, router.ResourceStatus(c.Request().Context()), nil)
}
