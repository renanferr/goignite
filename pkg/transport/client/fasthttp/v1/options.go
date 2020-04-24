package fasthttp

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Name                          string
	NoDefaultUserAgentHeader      bool          `config:"nodefaultuseragentheader"`
	MaxConnsPerHost               int           `config:"maxconnsperhost"`
	ReadBufferSize                int           `config:"readbuffersize"`
	WriteBufferSize               int           `config:"writebuffersize"`
	ReadTimeout                   time.Duration `config:"readtimeout"`
	WriteTimeout                  time.Duration `config:"writetimeout"`
	MaxIdleConnDuration           time.Duration `config:"maxidleconnduration"`
	DisableHeaderNamesNormalizing bool          `config:"disableheadernamesnormalizing"`
	DialDualStack                 bool          `config:"dialdualstack"`
	MaxResponseBodySize           int           `config:"maxresponsebodysize"`
	MaxIdemponentCallAttempts     int           `config:"maxidemponentcallattempts"`
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Name(value string) optionsBuilder {
	return builder.Set(b, "Name", value).(optionsBuilder)
}

func (b optionsBuilder) NoDefaultUserAgentHeader(value bool) optionsBuilder {
	return builder.Set(b, "NoDefaultUserAgentHeader", value).(optionsBuilder)
}

func (b optionsBuilder) MaxResponseBodySize(value int) optionsBuilder {
	return builder.Set(b, "MaxResponseBodySize", value).(optionsBuilder)
}

func (b optionsBuilder) MaxIdemponentCallAttempts(value int) optionsBuilder {
	return builder.Set(b, "MaxIdemponentCallAttempts", value).(optionsBuilder)
}

func (b optionsBuilder) MaxConnsPerHost(value int) optionsBuilder {
	return builder.Set(b, "MaxConnsPerHost", value).(optionsBuilder)
}

func (b optionsBuilder) WriteBufferSize(value int) optionsBuilder {
	return builder.Set(b, "WriteBufferSize", value).(optionsBuilder)
}

func (b optionsBuilder) ReadTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReadTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) WriteTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "WriteTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) MaxIdleConnDuration(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxIdleConnDuration", value).(optionsBuilder)
}

func (b optionsBuilder) DisableHeaderNamesNormalizing(value bool) optionsBuilder {
	return builder.Set(b, "DisableHeaderNamesNormalizing", value).(optionsBuilder)
}

func (b optionsBuilder) DialDualStack(value bool) optionsBuilder {
	return builder.Set(b, "DialDualStack", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.fasthttp", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
