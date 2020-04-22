package nats

import (
	"github.com/lann/builder"
)

type NewRelicOptions struct {
	Enabled bool
}

type newRelicOptionsBuilder builder.Builder

func (b newRelicOptionsBuilder) Enabled(value bool) newRelicOptionsBuilder {
	return builder.Set(b, "Enabled", value).(newRelicOptionsBuilder)
}

func (b newRelicOptionsBuilder) Build() NewRelicOptions {
	return builder.GetStruct(b).(NewRelicOptions)
}

var NewRelicOptionsBuilder = builder.Register(newRelicOptionsBuilder{}, NewRelicOptions{}).(newRelicOptionsBuilder)
