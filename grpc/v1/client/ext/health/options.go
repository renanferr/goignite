package gigrpchealth

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	"github.com/lann/builder"
)

type Options struct {
	Name        string
	Enabled     bool
	Description string
	Required    bool
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Name(name string) optionsBuilder {
	return builder.Set(b, "Name", name).(optionsBuilder)
}

func (b optionsBuilder) Enabled(enabled bool) optionsBuilder {
	return builder.Set(b, "Enabled", enabled).(optionsBuilder)
}

func (b optionsBuilder) Description(description string) optionsBuilder {
	return builder.Set(b, "Description", description).(optionsBuilder)
}

func (b optionsBuilder) Required(required bool) optionsBuilder {
	return builder.Set(b, "Required", required).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
