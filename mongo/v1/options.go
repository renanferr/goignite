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

func (b optionsBuilder) Uri(value string) optionsBuilder {
	return builder.Set(b, "Uri", value).(optionsBuilder)
}

func (b optionsBuilder) Auth(value *options.Credential) optionsBuilder {
	return builder.Set(b, "Auth", value).(optionsBuilder)
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
