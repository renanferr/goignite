package mongodb

import "github.com/lann/builder"

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
