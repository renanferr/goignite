package client

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	KubeConfigPath string
	Context        string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) KubeConfigPath(value string) optionsBuilder {
	return builder.Set(b, "KubeConfigPath", value).(optionsBuilder)
}

func (b optionsBuilder) Context(value string) optionsBuilder {
	return builder.Set(b, "Context", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
