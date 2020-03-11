package rest

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/health"
)

func Health(ctx context.Context) (HealthResponse, int) {

	var details []HealthDetailResponse

	all := health.CheckAll(ctx)

	httpStatus := http.StatusOK
	healthStatus := Ok

	for _, v := range all {

		healthDetailStatus := Ok

		if !v.IsOk() {
			healthDetailStatus = Down
		}

		var err string

		if v.Error != nil {
			err = v.Error.Error()
		}

		healthDetailResponse := HealthDetailResponseBuilder.
			Name(v.HealthCheck.Name).
			Description(v.HealthCheck.Description).
			Status(healthDetailStatus).
			Error(err).
			Build()

		details = append(details, healthDetailResponse)

		if !v.IsOk() && httpStatus != http.StatusServiceUnavailable {
			if v.HealthCheck.IsRequired() {
				httpStatus = http.StatusServiceUnavailable
				healthStatus = Down
			} else {
				httpStatus = http.StatusMultiStatus
				healthStatus = Partial
			}
		}
	}

	return HealthResponseBuilder.Details(details).Status(healthStatus).Build(), httpStatus
}
