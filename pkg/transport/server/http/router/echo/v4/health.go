package echo

import (
	"github.com/b2wdigital/goignite/pkg/transport/server/http/rest"
	"github.com/labstack/echo/v4"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {

	resp, httpCode := rest.Health(c.Request().Context())

	return JSONResponse(c, httpCode, resp, nil)
}
