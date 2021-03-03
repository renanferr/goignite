package status

import (
	"context"
	"net/http"

	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/rest/response"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	statusRoute := GetRoute()

	logger.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

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
