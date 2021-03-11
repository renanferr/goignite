package giechohealth

import (
	"context"

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

	healthRoute := GetRoute()

	logger.Tracef("configuring health router on %s in echo", healthRoute)

	healthHandler := NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)

	logger.Debugf("health router configured on %s in echo", healthRoute)

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
