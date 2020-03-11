package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
)

func NewConfig(ctx context.Context, options Options) aws.Config {

	l := log.FromContext(ctx)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		l.Panicf("unable to load AWS SDK config, %s", err.Error())
	}

	cfg.Region = options.DefaultRegion
	cfg.Credentials = aws.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)

	return cfg
}

func NewDefaultConfig(ctx context.Context) aws.Config {

	l := log.FromContext(ctx)

	o := Options{}

	err := config.UnmarshalWithPath("aws", &o)
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewConfig(ctx, o)

}
