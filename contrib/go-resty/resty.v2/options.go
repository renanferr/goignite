package resty

import (
	"time"

	"github.com/lann/builder"
)

type Options struct {
	Debug             bool
	ConnectionTimeout time.Duration
	KeepAlive         time.Duration
	RequestTimeout    time.Duration
	Transport         *OptionsTransport
	Host              string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) ConnectionTimeout(connTimeout time.Duration) optionsBuilder {
	return builder.Set(b, "ConnectionTimeout", connTimeout).(optionsBuilder)
}

func (b optionsBuilder) KeepAlive(keepalive time.Duration) optionsBuilder {
	return builder.Set(b, "KeepAlive", keepalive).(optionsBuilder)
}

func (b optionsBuilder) RequestTimeout(timeout time.Duration) optionsBuilder {
	return builder.Set(b, "RequestTimeout", timeout).(optionsBuilder)
}

func (b optionsBuilder) Host(host string) optionsBuilder {
	return builder.Set(b, "Host", host).(optionsBuilder)
}

func (b optionsBuilder) Debug(debug bool) optionsBuilder {
	return builder.Set(b, "Debug", debug).(optionsBuilder)
}

func (b optionsBuilder) Transport(transport *OptionsTransport) optionsBuilder {
	return builder.Set(b, "Transport", transport).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)
