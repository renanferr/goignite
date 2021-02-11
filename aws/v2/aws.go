package giaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	giconfig "github.com/b2wdigital/goignite/config"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
)

const (
	TopicConfig = "topic:giaws:config"
)

func NewConfig(ctx context.Context, options *Options) aws.Config {

	l := gilog.FromContext(ctx)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		l.Panicf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}
	}

	cfg.Region = options.DefaultRegion

	if options.SessionToken == "" {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)
	}

	gieventbus.Publish(TopicConfig, &cfg)

	return cfg
}

func NewDefaultConfig(ctx context.Context) aws.Config {

	o := loadDefaultOptions(ctx)

	return NewConfig(ctx, o)
}

func loadDefaultOptions(ctx context.Context) *Options {

	logger := gilog.FromContext(ctx)

	o := &Options{}

	var err error

	err = giconfig.UnmarshalWithPath("aws.access.key", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = giconfig.UnmarshalWithPath("aws.secret.access", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = giconfig.UnmarshalWithPath("aws.default", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = giconfig.UnmarshalWithPath("aws.session", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return o
}
