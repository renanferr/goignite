package newrelic

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Enabled bool
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Enabled(value bool) optionsBuilder {
	return builder.Set(b, "Enabled", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("aws.integration.newrelic", o)
	if err != nil {
		return nil, err
	}

	return o, nil

}
