package nats

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Url           string
	MaxReconnects int           `config:"maxreconnects"`
	ReconnectWait time.Duration `config:"reconnectwait"`
	Health        OptionsHealth
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Url(value string) optionsBuilder {
	return builder.Set(b, "Url", value).(optionsBuilder)
}

func (b optionsBuilder) MaxReconnects(value int) optionsBuilder {
	return builder.Set(b, "MaxReconnects", value).(optionsBuilder)
}

func (b optionsBuilder) ReconnectWait(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReconnectWait", value).(optionsBuilder)
}

func (b optionsBuilder) Health(health OptionsHealth) optionsBuilder {
	return builder.Set(b, "Health", health).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.nats", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
