package gigocloud

import (
	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/lann/builder"
)

// Options ..
type Options struct {
	Resource string
	Type     string
	Region   string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Resource(resource string) optionsBuilder {
	return builder.Set(b, "Resource", resource).(optionsBuilder)
}

func (b optionsBuilder) Type(tp string) optionsBuilder {
	return builder.Set(b, "Type", tp).(optionsBuilder)
}

func (b optionsBuilder) Region(region string) optionsBuilder {
	return builder.Set(b, "Region", region).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

// OptionsBuilder ..
var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil

}
