package router

import "github.com/lann/builder"

type HealthDetailResponse struct {
	Status      HealthStatus `json:"status" binding:"required"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description,omitempty" binding:"required"`
	Error       string       `json:"error,omitempty"`
}

type healthDetailResponseBuilder builder.Builder

func (b healthDetailResponseBuilder) Status(value HealthStatus) healthDetailResponseBuilder {
	return builder.Set(b, "Status", value).(healthDetailResponseBuilder)
}

func (b healthDetailResponseBuilder) Name(value string) healthDetailResponseBuilder {
	return builder.Set(b, "Name", value).(healthDetailResponseBuilder)
}

func (b healthDetailResponseBuilder) Description(value string) healthDetailResponseBuilder {
	return builder.Set(b, "Description", value).(healthDetailResponseBuilder)
}

func (b healthDetailResponseBuilder) Error(value string) healthDetailResponseBuilder {
	return builder.Set(b, "Error", value).(healthDetailResponseBuilder)
}

func (b healthDetailResponseBuilder) Build() HealthDetailResponse {
	return builder.GetStruct(b).(HealthDetailResponse)
}

var HealthDetailResponseBuilder = builder.Register(healthDetailResponseBuilder{}, HealthDetailResponse{}).(healthDetailResponseBuilder)
