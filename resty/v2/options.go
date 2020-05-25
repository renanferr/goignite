package giresty

import (
	"time"

	"github.com/lann/builder"
)

type Options struct {
	Debug          bool
	RequestTimeout time.Duration
	Retry          OptionsRetry
	Host           string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) RequestTimeout(timeout time.Duration) optionsBuilder {
	return builder.Set(b, "RequestTimeout", timeout).(optionsBuilder)
}

func (b optionsBuilder) Retry(retry OptionsRetry) optionsBuilder {
	return builder.Set(b, "Retry", retry).(optionsBuilder)
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
