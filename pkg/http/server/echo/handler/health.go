package handler

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/health"
	"github.com/labstack/echo"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {

	all := health.CheckAll(c.Request().Context())

	status := http.StatusOK

	for _, v := range all {
		if !v.IsOk() && status != http.StatusServiceUnavailable {
			if v.HealthCheck.IsRequired() {
				status = http.StatusServiceUnavailable
				break
			} else {
				status = http.StatusMultiStatus
			}
		}
	}

	return c.JSON(status, all)
}
