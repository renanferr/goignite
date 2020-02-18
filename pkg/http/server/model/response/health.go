package response

import "github.com/lann/builder"

type HealthResponse struct {
	Status  HealthStatus           `json:"status" binding:"required"`
	Details []HealthDetailResponse `json:"details" binding:"required"`
}

type healthResponseBuilder builder.Builder

func (b healthResponseBuilder) Status(value HealthStatus) healthResponseBuilder {
	return builder.Set(b, "Status", value).(healthResponseBuilder)
}

func (b healthResponseBuilder) Details(value []HealthDetailResponse) healthResponseBuilder {
	return builder.Set(b, "Details", value).(healthResponseBuilder)
}

func (b healthResponseBuilder) Build() HealthResponse {
	return builder.GetStruct(b).(HealthResponse)
}

var HealthResponseBuilder = builder.Register(healthResponseBuilder{}, HealthResponse{}).(healthResponseBuilder)
