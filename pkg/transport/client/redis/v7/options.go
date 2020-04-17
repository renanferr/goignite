package redis

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Password           string
	MaxRetries         int           `config:"maxretries"`
	MinRetryBackoff    time.Duration `config:"minretrybackoff"`
	MaxRetryBackoff    time.Duration `config:"maxretrybackoff"`
	DialTimeout        time.Duration `config:"dialtimeout"`
	ReadTimeout        time.Duration `config:"readtimeout"`
	WriteTimeout       time.Duration `config:"writetimeout"`
	PoolSize           int           `config:"poolsize"`
	MinIdleConns       int           `config:"minidleconns"`
	MaxConnAge         time.Duration `config:"maxconnage"`
	PoolTimeout        time.Duration `config:"pooltimeout"`
	IdleTimeout        time.Duration `config:"idletimeout"`
	IdleCheckFrequency time.Duration `config:"idlecheckfrequency"`
	Client             ClientOptions
	Cluster            ClusterOptions
	Health             HealthOptions
	NewRelic           NewRelicOptions
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Password(value string) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRetries(value int) optionsBuilder {
	return builder.Set(b, "MaxRetries", value).(optionsBuilder)
}

func (b optionsBuilder) MinRetryBackoff(value time.Duration) optionsBuilder {
	return builder.Set(b, "MinRetryBackoff", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRetryBackoff(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxRetryBackoff", value).(optionsBuilder)
}

func (b optionsBuilder) ReadTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReadTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) DialTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "DialTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) WriteTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "WriteTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) PoolSize(value int) optionsBuilder {
	return builder.Set(b, "PoolSize", value).(optionsBuilder)
}

func (b optionsBuilder) MinIdleConns(value int) optionsBuilder {
	return builder.Set(b, "MinIdleConns", value).(optionsBuilder)
}

func (b optionsBuilder) MaxConnAge(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxConnAge", value).(optionsBuilder)
}

func (b optionsBuilder) PoolTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "PoolTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) IdleTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "IdleTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) IdleCheckFrequency(value time.Duration) optionsBuilder {
	return builder.Set(b, "IdleCheckFrequency", value).(optionsBuilder)
}

func (b optionsBuilder) Health(value HealthOptions) optionsBuilder {
	return builder.Set(b, "Health", value).(optionsBuilder)
}

func (b optionsBuilder) NewRelic(value NewRelicOptions) optionsBuilder {
	return builder.Set(b, "NewRelic", value).(optionsBuilder)
}

func (b optionsBuilder) Client(value ClientOptions) optionsBuilder {
	return builder.Set(b, "Client", value).(optionsBuilder)
}

func (b optionsBuilder) Cluster(value ClusterOptions) optionsBuilder {
	return builder.Set(b, "Cluster", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.redis", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
