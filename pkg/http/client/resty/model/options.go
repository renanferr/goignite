package model

import (
	"github.com/lann/builder"
)

type Options struct {
	Debug          bool
	RequestTimeout int
	Retry          OptionsRetry
	Health         OptionsHealth
	Host           string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) RequestTimeout(timeout int) optionsBuilder {
	return builder.Set(b, "RequestTimeout", timeout).(optionsBuilder)
}

func (b optionsBuilder) Retry(retry OptionsRetry) optionsBuilder {
	return builder.Set(b, "Retry", retry).(optionsBuilder)
}

func (b optionsBuilder) Health(health OptionsHealth) optionsBuilder {
	return builder.Set(b, "Health", health).(optionsBuilder)
}

func (b optionsBuilder) Host(host string) optionsBuilder {
	return builder.Set(b, "Host", host).(optionsBuilder)
}

func (b optionsBuilder) Debug(debug bool) optionsBuilder {
	return builder.Set(b, "Debug", debug).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)
