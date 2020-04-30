package mongodb

import (
	"github.com/lann/builder"
)

type OptionsNewRelic struct {
	Enabled bool
}

type optionsNewRelicBuilder builder.Builder

func (b optionsNewRelicBuilder) Enabled(value bool) optionsNewRelicBuilder {
	return builder.Set(b, "Enabled", value).(optionsNewRelicBuilder)
}

func (b optionsNewRelicBuilder) Build() OptionsNewRelic {
	return builder.GetStruct(b).(OptionsNewRelic)
}

var OptionsNewRelicBuilder = builder.Register(optionsNewRelicBuilder{}, OptionsNewRelic{}).(optionsNewRelicBuilder)
