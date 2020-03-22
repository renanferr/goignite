package elasticsearch

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Addresses             []string
	Username              string
	Password              string
	CloudID               string        `config:"cloudid"`
	APIKey                string        `config:"apikey"`
	CACert                string        `config:"cacert"`
	RetryOnStatus         []string      `config:"retryonstatus"`
	DisableRetry          bool          `config:"disableretry"`
	EnableRetryOnTimeout  bool          `config:"enableretryontimeout"`
	MaxRetries            int           `config:"maxretries"`
	DiscoverNodesOnStart  bool          `config:"discovernodesonstart"`
	DiscoverNodesInterval time.Duration `config:"discovernodesinterval"`
	EnableMetrics         bool          `config:"enablemetrics"`
	EnableDebugLogger     bool          `config:"enabledebuglogger"`
	RetryBackoff          time.Duration `config:"retrybackoff"`
	Health                OptionsHealth
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Addresses(value []string) optionsBuilder {
	return builder.Set(b, "Addresses", value).(optionsBuilder)
}

func (b optionsBuilder) Username(value string) optionsBuilder {
	return builder.Set(b, "Username", value).(optionsBuilder)
}

func (b optionsBuilder) Password(value string) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) CloudID(value string) optionsBuilder {
	return builder.Set(b, "CloudID", value).(optionsBuilder)
}

func (b optionsBuilder) APIKey(value string) optionsBuilder {
	return builder.Set(b, "APIKey", value).(optionsBuilder)
}

func (b optionsBuilder) CACert(value string) optionsBuilder {
	return builder.Set(b, "CACert", value).(optionsBuilder)
}

func (b optionsBuilder) RetryOnStatus(value []string) optionsBuilder {
	return builder.Set(b, "RetryOnStatus", value).(optionsBuilder)
}

func (b optionsBuilder) DisableRetry(value bool) optionsBuilder {
	return builder.Set(b, "DisableRetry", value).(optionsBuilder)
}

func (b optionsBuilder) EnableRetryOnTimeout(value bool) optionsBuilder {
	return builder.Set(b, "EnableRetryOnTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRetries(value int) optionsBuilder {
	return builder.Set(b, "MaxRetries", value).(optionsBuilder)
}

func (b optionsBuilder) DiscoverNodesOnStart(value bool) optionsBuilder {
	return builder.Set(b, "DiscoverNodesOnStart", value).(optionsBuilder)
}

func (b optionsBuilder) DiscoverNodesInterval(value time.Duration) optionsBuilder {
	return builder.Set(b, "DiscoverNodesInterval", value).(optionsBuilder)
}

func (b optionsBuilder) EnableMetrics(value bool) optionsBuilder {
	return builder.Set(b, "EnableMetrics", value).(optionsBuilder)
}

func (b optionsBuilder) EnableDebugLogger(value bool) optionsBuilder {
	return builder.Set(b, "EnableDebugLogger", value).(optionsBuilder)
}

func (b optionsBuilder) RetryBackoff(value time.Duration) optionsBuilder {
	return builder.Set(b, "RetryBackoff", value).(optionsBuilder)
}

func (b optionsBuilder) Health(health OptionsHealth) optionsBuilder {
	return builder.Set(b, "Health", health).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)


func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.elasticsearch", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}