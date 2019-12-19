package model

import "github.com/lann/builder"

type Options struct {
	Debug            bool
	RequestTimeout   int
	RetryCount       int
	RetryWaitTime    int
	RetryMaxWaitTime int
	Host             string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) RequestTimeout(timeout int) optionsBuilder {
	return builder.Set(b, "RequestTimeout", timeout).(optionsBuilder)
}

func (b optionsBuilder) RetryCount(retryCount int) optionsBuilder {
	return builder.Set(b, "RetryCount", retryCount).(optionsBuilder)
}

func (b optionsBuilder) RetryWaitTime(retryWaitTime int) optionsBuilder {
	return builder.Set(b, "RetryWaitTime", retryWaitTime).(optionsBuilder)
}

func (b optionsBuilder) RetryMaxWaitTime(retryMaxWaitTime int) optionsBuilder {
	return builder.Set(b, "RetryMaxWaitTime", retryMaxWaitTime).(optionsBuilder)
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
