package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	localaws "github.com/b2wdigital/goignite/pkg/cloud/aws/v2"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	_ "github.com/b2wdigital/goignite/pkg/log/zap/v1"
)

const Bucket = "cloud.aws.s3.bucket"

func init() {
	config.Add(Bucket, "example", "s3 example bucket")
}

func main() {

	// parse config
	var err error

	config.Load()

	// create background context
	ctx := context.Background()

	// start logrus
// 	log.NewLogger(zap.NewLogger())

	// get logrus instance from context
	l := log.FromContext(ctx)

	// create default aws config
	awsConfig := localaws.NewDefaultConfig(ctx)

	// create s3 client
	s3Client := s3.New(awsConfig)

	// set vars
	filename := "examplefile"
	bucket := config.String(Bucket)

	// prepare s3 request head
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}

	// make a call
	req := s3Client.HeadObjectRequest(input)

	head, err := req.Send(ctx)

	if err != nil {

		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			log.Fatalf(err.Error())
		}

		l.Fatalf("unable check file %s in s3 bucket %s", filename, bucket)
	}

	l = l.WithFields(
		log.Fields{"lastModified": head.LastModified,
			"versionId": head.VersionId,
		})

	l.Debugf("file %s exists on bucket %s", filename, bucket)

}
