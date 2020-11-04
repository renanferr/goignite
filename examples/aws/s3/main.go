package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	giaws "github.com/b2wdigital/goignite/aws/v2"
	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"github.com/prometheus/common/log"
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
	l := gilog.FromContext(ctx)

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
		log.Fatalf(err.Error())
	}

	if err != nil {

		l.Fatalf("unable check file %s in s3 bucket %s", filename, bucket)
	}

	l = l.WithFields(
		gilog.Fields{"lastModified": head.LastModified,
			"versionId": head.VersionId,
		})

	l.Debugf("file %s exists on bucket %s", filename, bucket)

}
