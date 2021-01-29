package giresty

import (
	"time"

	"github.com/lann/builder"
)

type Options struct {
	Debug             bool
	ConnectionTimeout time.Duration
	KeepAlive         time.Duration
	RequestTimeout    time.Duration
	Retry             *OptionsRetry
	Transport         *OptionsTransport
	Host              string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) ConnectionTimeout(connTimeout time.Duration) optionsBuilder {
	return builder.Set(b, "connectionTimeout", connTimeout).(optionsBuilder)
}

func (b optionsBuilder) KeepAlive(keepalive time.Duration) optionsBuilder {
	return builder.Set(b, "keepAlive", keepalive).(optionsBuilder)
}

func (b optionsBuilder) RequestTimeout(timeout time.Duration) optionsBuilder {
	return builder.Set(b, "requestTimeout", timeout).(optionsBuilder)
}

func (b optionsBuilder) Retry(retry *OptionsRetry) optionsBuilder {
	return builder.Set(b, "Retry", retry).(optionsBuilder)
}

func (b optionsBuilder) Host(host string) optionsBuilder {
	return builder.Set(b, "Host", host).(optionsBuilder)
}

func (b optionsBuilder) Debug(debug bool) optionsBuilder {
	return builder.Set(b, "debug", debug).(optionsBuilder)
}

func (b optionsBuilder) Transport(transport *OptionsTransport) optionsBuilder {
	return builder.Set(b, "Transport", transport).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)
