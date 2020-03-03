package service

import (
	"context"
	"net/http"

	"github.com/jpfaria/goignite/pkg/health"
	"github.com/jpfaria/goignite/pkg/http/server/model/response"
)

func Health(ctx context.Context) response.HealthResponse {

	var details []response.HealthDetailResponse

	all := health.CheckAll(ctx)

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

	return response.HealthResponseBuilder.Details(details).Status(healthStatus).Build()
}
