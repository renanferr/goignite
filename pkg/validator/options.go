package validator

import (
	"github.com/lann/builder"
)

type Options struct {
	Locale string
	Schema map[string]string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) SetLocale(locale string) optionsBuilder {
	return builder.Set(b, "Locale", locale).(optionsBuilder)
}

func (b optionsBuilder) SetSchema(schema map[string]string) optionsBuilder {
	return builder.Set(b, "Schema", schema).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)
