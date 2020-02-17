package model

import "github.com/lann/builder"

type OptionsRetry struct {
	Count       int
	WaitTime    int
	MaxWaitTime int
}

type optionsRetryBuilder builder.Builder

func (b optionsRetryBuilder) RetryCount(retryCount int) optionsRetryBuilder {
	return builder.Set(b, "Count", retryCount).(optionsRetryBuilder)
}

func (b optionsRetryBuilder) RetryWaitTime(retryWaitTime int) optionsRetryBuilder {
	return builder.Set(b, "WaitTime", retryWaitTime).(optionsRetryBuilder)
}

func (b optionsRetryBuilder) RetryMaxWaitTime(retryMaxWaitTime int) optionsRetryBuilder {
	return builder.Set(b, "MaxWaitTime", retryMaxWaitTime).(optionsRetryBuilder)
}

func (b optionsRetryBuilder) Build() OptionsRetry {
	return builder.GetStruct(b).(OptionsRetry)
}

var OptionsRetryBuilder = builder.Register(optionsRetryBuilder{}, Options{}).(optionsRetryBuilder)
