package godror

import "github.com/lann/builder"

type OptionsHealth struct {
	Enabled     bool
	Description string
	Required    bool
}

type optionsHealthBuilder builder.Builder

func (b optionsHealthBuilder) Enabled(enabled bool) optionsHealthBuilder {
	return builder.Set(b, "Enabled", enabled).(optionsHealthBuilder)
}

func (b optionsHealthBuilder) Description(description string) optionsHealthBuilder {
	return builder.Set(b, "Description", description).(optionsHealthBuilder)
}

func (b optionsHealthBuilder) Required(required bool) optionsHealthBuilder {
	return builder.Set(b, "Required", required).(optionsHealthBuilder)
}

func (b optionsHealthBuilder) Build() OptionsHealth {
	return builder.GetStruct(b).(OptionsHealth)
}

var OptionsHealthBuilder = builder.Register(optionsHealthBuilder{}, OptionsHealth{}).(optionsHealthBuilder)
