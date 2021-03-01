package health

import (
	"context"

	giecho "github.com/b2wdigital/goignite/echo/v4"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/labstack/echo/v4"
)

func Route(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {

		logger := gilog.FromContext(ctx)

		healthRoute := getRoute()

		logger.Infof("configuring health router on %s", healthRoute)

		healthHandler := NewHealthHandler()
		instance.GET(healthRoute, healthHandler.Get)

	}

	return nil
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	resp, httpCode := response.NewHealth(ctx)

	return giecho.JSON(c, httpCode, resp, nil)
}
