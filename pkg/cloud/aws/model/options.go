package model

import "github.com/lann/builder"

type Options struct {
	AccessKeyId     string `config:"access.key.id"`
	SecretAccessKey string `config:"secret.access.key"`
	DefaultRegion   string `config:"default.region"`
	SessionToken    string `config:"session.token"`
}

type optionsBuilder builder.Builder

func (b optionsBuilder) AccessKeyId(value string) optionsBuilder {
	return builder.Set(b, "AccessKeyId", value).(optionsBuilder)
}

func (b optionsBuilder) SecretAccessKey(value string) optionsBuilder {
	return builder.Set(b, "AccessKeSecretAccessKeyyId", value).(optionsBuilder)
}

func (b optionsBuilder) DefaultRegion(value string) optionsBuilder {
	return builder.Set(b, "DefaultRegion", value).(optionsBuilder)
}

func (b optionsBuilder) SessionToken(value string) optionsBuilder {
	return builder.Set(b, "SessionToken", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}
