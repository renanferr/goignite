package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
)

func NewConfigWithIntegrations(ctx context.Context, options *Options, integrations []Integrator) aws.Config {

	cfg := NewConfig(ctx, options)

	for _, integrator := range integrations {
		err := integrator.Integrate(ctx, &cfg)
		if err != nil {
			continue
		}
	}

	return cfg
}

func NewConfig(ctx context.Context, options *Options) aws.Config {

	l := log.FromContext(ctx)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		l.Panicf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}
	}

	cfg.Region = options.DefaultRegion

	if options.SessionToken == "" {
		cfg.Credentials = aws.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)
	}

	return cfg
}

func NewDefaultConfig(ctx context.Context) aws.Config {

	o := loadDefaultOptions(ctx)

	return NewConfig(ctx, o)
}

func NewDefaultConfigWithIntegrations(ctx context.Context, integrations []Integrator) aws.Config {

	o := loadDefaultOptions(ctx)

	return NewConfigWithIntegrations(ctx, o, integrations)
}

func loadDefaultOptions(ctx context.Context) *Options {

	logger := log.FromContext(ctx)

	o := &Options{}

	var err error

	err = config.UnmarshalWithPath("aws.access.key", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = config.UnmarshalWithPath("aws.secret.access", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = config.UnmarshalWithPath("aws.default", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = config.UnmarshalWithPath("aws.session", o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return o
}
