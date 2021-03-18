package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/http/v1/client"
	"github.com/b2wdigital/goignite/v2/log"
)

type Ext func(context.Context, *aws.Config) error

func NewConfig(ctx context.Context, options *Options, exts ...Ext) aws.Config {

	logger := log.FromContext(ctx)

	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Panicf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}
	}

	cfg.Region = options.DefaultRegion

	if options.SessionToken == "" {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)
	}

	httpClientOptions := client.Options{}

	err = config.UnmarshalWithPath(httpClientRoot, &httpClientOptions)
	if err != nil {
		logger.Panicf("unable to load http client config, %s", err.Error())
		return aws.Config{}
	}

	httpClient := client.NewClient(ctx, &httpClientOptions)

	cfg.Retryer = retryerConfig(options)
	cfg.HTTPClient = httpClient

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

	err = config.UnmarshalWithPath(retryerRoot, o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = config.UnmarshalWithPath(httpClientRoot, o)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return o
}
