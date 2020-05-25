package sender

import (
	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/lann/builder"
)

type Options struct {
	Url      string
	Subjects []string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Url(value string) optionsBuilder {
	return builder.Set(b, "Url", value).(optionsBuilder)
}

func (b optionsBuilder) Subjects(value []string) optionsBuilder {
	return builder.Set(b, "Subjects", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath("gi.cloudevents.nats.sender", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
