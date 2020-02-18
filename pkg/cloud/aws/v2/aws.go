package v2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
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

	if options.CustomEndpoint != "" {
		cfg.EndpointResolver = aws.EndpointResolverFunc(endpointResolver(cfg.Region))
	}

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

func endpointResolver(awsRegion string) func(service, region string) (aws.Endpoint, error) {
	return func(service, region string) (aws.Endpoint, error) {
		if service == kinesis.EndpointsID {
			return aws.Endpoint{
				URL:           "http://localhost:4568",
				SigningRegion: awsRegion,
			}, nil
		}

		if service == s3.EndpointsID {
			return aws.Endpoint{
				URL:           "http://localhost:4572",
				SigningRegion: awsRegion,
			}, nil
		}

		if service == sqs.EndpointsID {
			return aws.Endpoint{
				URL:           "http://localhost:4576",
				SigningRegion: awsRegion,
			}, nil
		}

		if service == sns.EndpointsID {
			return aws.Endpoint{
				URL:           "http://localhost:4575",
				SigningRegion: awsRegion,
			}, nil
		}

		return endpoints.NewDefaultResolver().ResolveEndpoint(service, region)
	}
}