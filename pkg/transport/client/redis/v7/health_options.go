package redis

import "github.com/lann/builder"

type HealthOptions struct {
	Enabled     bool
	Description string
	Required    bool
}

type healthOptionsBuilder builder.Builder

func (b healthOptionsBuilder) Enabled(enabled bool) healthOptionsBuilder {
	return builder.Set(b, "Enabled", enabled).(healthOptionsBuilder)
}

func (b healthOptionsBuilder) Description(description string) healthOptionsBuilder {
	return builder.Set(b, "Description", description).(healthOptionsBuilder)
}

func (b healthOptionsBuilder) Required(required bool) healthOptionsBuilder {
	return builder.Set(b, "Required", required).(healthOptionsBuilder)
}

func (b healthOptionsBuilder) Build() HealthOptions {
	return builder.GetStruct(b).(HealthOptions)
}

var HealthOptionsBuilder = builder.Register(healthOptionsBuilder{}, HealthOptions{}).(healthOptionsBuilder)
