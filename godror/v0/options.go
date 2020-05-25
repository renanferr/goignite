package gigodror

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/lann/builder"
)

type Options struct {
	DataSourceName  string        `config:"datasourcename"`
	ConnMaxLifetime time.Duration `config:"connmaxlifetime"`
	MaxIdleConns    int           `config:"maxidleconns"`
	MaxOpenConns    int           `config:"maxopenconns"`
	Health          OptionsHealth
}

type optionsBuilder builder.Builder

func (b optionsBuilder) DataSourceName(value string) optionsBuilder {
	return builder.Set(b, "DataSourceName", value).(optionsBuilder)
}

func (b optionsBuilder) ConnMaxLifetime(value time.Duration) optionsBuilder {
	return builder.Set(b, "ConnMaxLifetime", value).(optionsBuilder)
}

func (b optionsBuilder) MaxIdleConns(value int) optionsBuilder {
	return builder.Set(b, "MaxIdleConns", value).(optionsBuilder)
}

func (b optionsBuilder) MaxOpenConns(value int) optionsBuilder {
	return builder.Set(b, "MaxOpenConns", value).(optionsBuilder)
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

	err := giconfig.UnmarshalWithPath("gi.godror", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
