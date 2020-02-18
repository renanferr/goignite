package handler

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/health"
	"github.com/jpfaria/goignite/pkg/http/server/model/response"
	"github.com/labstack/echo"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {

	var details []response.HealthDetailResponse

	all := health.CheckAll(c.Request().Context())

	httpStatus := http.StatusOK
	healthStatus := response.Ok


	for _, v := range all {

		healthDetailStatus := response.Ok

		if !v.IsOk() {
			healthDetailStatus = response.Down
		}

		var err string

		if v.Error != nil {
			err = v.Error.Error()
		}

		healthDetailResponse := response.HealthDetailResponseBuilder.
			Name(v.HealthCheck.Name).
			Description(v.HealthCheck.Description).
			Status(healthDetailStatus).
			Error(err).
			Build()


		details = append(details, healthDetailResponse)

		if !v.IsOk() && httpStatus != http.StatusServiceUnavailable {
			if v.HealthCheck.IsRequired() {
				httpStatus = http.StatusServiceUnavailable
				healthStatus = response.Down
			} else {
				httpStatus = http.StatusMultiStatus
				healthStatus = response.Partial
			}
		}
	}

	healthResponse := response.HealthResponseBuilder.Details(details).Status(healthStatus).Build()

	return c.JSON(httpStatus, healthResponse)
}
