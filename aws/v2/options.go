package giaws

import "github.com/lann/builder"

type Options struct {
	AccessKeyId     string `config:"id"`
	SecretAccessKey string `config:"key"`
	DefaultRegion   string `config:"region"`
	SessionToken    string `config:"token"`
	MaxAttempts     int    `config:"maxattempts"`
	HasRateLimit    bool   `config:"hasratelimit"`
}

type optionsBuilder builder.Builder

func (b optionsBuilder) AccessKeyId(value string) optionsBuilder {
	return builder.Set(b, "AccessKeyId", value).(optionsBuilder)
}

func (b optionsBuilder) SecretAccessKey(value string) optionsBuilder {
	return builder.Set(b, "SecretAccessKey", value).(optionsBuilder)
}

func (b optionsBuilder) DefaultRegion(value string) optionsBuilder {
	return builder.Set(b, "DefaultRegion", value).(optionsBuilder)
}

func (b optionsBuilder) SessionToken(value string) optionsBuilder {
	return builder.Set(b, "SessionToken", value).(optionsBuilder)
}

func (b optionsBuilder) MaxAttempts(value int) optionsBuilder {
	return builder.Set(b, "MaxAttempts", value).(optionsBuilder)
}

func (b optionsBuilder) HasRateLimit(value bool) optionsBuilder {
	return builder.Set(b, "HasRateLimit", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)
