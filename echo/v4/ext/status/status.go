package status

import (
	"context"
	"net/http"

	giecho "github.com/b2wdigital/goignite/echo/v4"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/labstack/echo/v4"
)

func Route(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {

		logger := gilog.FromContext(ctx)

		statusRoute := getRoute()

		logger.Infof("configuring status router on %s", statusRoute)

		statusHandler := NewResourceStatusHandler()
		instance.GET(statusRoute, statusHandler.Get)
	}

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
