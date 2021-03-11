package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	giaws "github.com/b2wdigital/goignite/v2/aws/v2"
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
)

const Bucket = "aws.s3.bucket"

func init() {
	giconfig.Add(Bucket, "example", "s3 example bucket")
}

func main() {

	giconfig.Load()

	// create background context
	ctx := context.Background()

	// start logrus
	// zap.NewLogger()
	gilogrus.NewLogger()

	// get logrus instance from context
	logger := gilog.FromContext(ctx)

	// create default aws config
	awsConfig := giaws.NewDefaultConfig(ctx)

	// create s3 client

	s3Client := s3.NewFromConfig(awsConfig)

	// set vars
	filename := "examplefile"
	bucket := giconfig.String(Bucket)

	// prepare s3 request head
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}

	// make a call
	head, err := s3Client.HeadObject(ctx, input)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if err != nil {

		logger.Fatalf("unable check file %s in s3 bucket %s", filename, bucket)
	}

	logger = logger.WithFields(
		gilog.Fields{"lastModified": head.LastModified,
			"versionId": head.VersionId,
		})

	logger.Debugf("file %s exists on bucket %s", filename, bucket)

}
