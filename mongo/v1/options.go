package gimongo

import (
	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/lann/builder"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	Uri  string
	Auth *options.Credential
}

type optionsBuilder builder.Builder

func (b optionsBuilder) RequestTimeout(uri string) optionsBuilder {
	return builder.Set(b, "Uri", uri).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(ConfigRoot, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
