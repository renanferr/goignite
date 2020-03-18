package mongodb

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Uri    string
	Health OptionsHealth
}

type optionsBuilder builder.Builder

func (b optionsBuilder) RequestTimeout(uri string) optionsBuilder {
	return builder.Set(b, "Uri", uri).(optionsBuilder)
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

	err := config.UnmarshalWithPath("transport.client.mongodb", o)
	if err != nil {
		return nil, err
	}

	return o, nil

}
