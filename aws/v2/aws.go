package giaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
)

type Ext func(context.Context, *aws.Config) error

func NewConfig(ctx context.Context, options *Options, exts ...Ext) aws.Config {

	logger := gilog.FromContext(ctx)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Panicf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}
	}

	cfg.Region = options.DefaultRegion

	if options.SessionToken == "" {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)
	}

	cfg.Retryer = retryerConfig(options)

	for _, ext := range exts {
		if err := ext(ctx, &cfg); err != nil {
			panic(err)
		}
	}

	return cfg
}

func retryerConfig(options *Options) func() aws.Retryer {

	return func() aws.Retryer {

		return retry.NewStandard(func(o *retry.StandardOptions) {

			o.MaxAttempts = options.MaxAttempts

			if !options.HasRateLimit {
				o.RateLimiter = noRateLimit{}
			}

		})
	}
}

type noRateLimit struct{}

func (noRateLimit) AddTokens(uint) error { return nil }

func (noRateLimit) GetToken(context.Context, uint) (func() error, error) { return nil, nil }

func NewDefaultConfig(ctx context.Context, exts ...Ext) aws.Config {

	o := loadDefaultOptions(ctx)

	return NewConfig(ctx, o, exts...)
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

	err = giconfig.UnmarshalWithPath(retryer, o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return o
}
