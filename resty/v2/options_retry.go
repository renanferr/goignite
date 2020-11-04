package giresty

import (
	"time"

	"github.com/lann/builder"
)

type OptionsRetry struct {
	Count       int
	WaitTime    time.Duration
	MaxWaitTime time.Duration
}

type optionsRetryBuilder builder.Builder

func (b optionsRetryBuilder) RetryCount(retryCount int) optionsRetryBuilder {
	return builder.Set(b, "Count", retryCount).(optionsRetryBuilder)
}

func (b optionsRetryBuilder) RetryWaitTime(retryWaitTime time.Duration) optionsRetryBuilder {
	return builder.Set(b, "WaitTime", retryWaitTime).(optionsRetryBuilder)
}

func (b optionsRetryBuilder) RetryMaxWaitTime(retryMaxWaitTime time.Duration) optionsRetryBuilder {
	return builder.Set(b, "MaxWaitTime", retryMaxWaitTime).(optionsRetryBuilder)
}

func (b optionsRetryBuilder) Build() OptionsRetry {
	return builder.GetStruct(b).(OptionsRetry)
}

var OptionsRetryBuilder = builder.Register(optionsRetryBuilder{}, Options{}).(optionsRetryBuilder)
