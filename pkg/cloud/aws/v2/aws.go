package v2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/jpfaria/goignite/pkg/cloud/aws/model"
	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/logging/logrus"
)

func NewConfig(ctx context.Context, options model.Options) aws.Config {

	log := logrus.FromContext(ctx)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Panicf("unable to load AWS SDK config, %s", err.Error())
	}

	cfg.Region = options.DefaultRegion
	cfg.Credentials = aws.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)

	return cfg
}

func NewDefaultConfig(ctx context.Context) aws.Config {

	log := logrus.FromContext(ctx)

	o := model.Options{}

	err := config.UnmarshalWithPath("cloud.aws", &o)
	if err != nil {
		log.Fatal(err)
	}

	return NewConfig(ctx, o)

}