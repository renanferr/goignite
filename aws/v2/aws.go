package aws

import (
	"context"
	"net"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/b2wdigital/goignite/v2/config"
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

	cfg.Retryer = retryerConfig(options)
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: options.MaxConnsPerHost,
			Proxy:           http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   options.TimeoutMillis,
				KeepAlive: options.KeepAliveMillis,
			}).DialContext,
			MaxIdleConns:          options.MaxIdleConns,
			MaxIdleConnsPerHost:   options.MaxIdleConnsPerHost,
			ResponseHeaderTimeout: options.ResponseHeaderTimeoutMillis,
			IdleConnTimeout:       options.IdleConnTimeoutMillis,
		},
	}

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
