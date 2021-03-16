package giechostatus

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/rest/response"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	statusRoute := GetRoute()

	logger.Tracef("configuring status router on %s in echo", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	logger.Debugf("status router configured on %s in echo", statusRoute)

	return nil
}

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c echo.Context) error {
	return giecho.JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
}
