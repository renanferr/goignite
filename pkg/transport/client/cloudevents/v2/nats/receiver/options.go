package receiver

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Subjects []string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Subjects(value []string) optionsBuilder {
	return builder.Set(b, "Subjects", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.cloudevents.nats.receiver", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
