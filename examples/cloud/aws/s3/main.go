package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	v2 "github.com/jpfaria/goignite/pkg/cloud/aws/v2"
	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/log/logrus"
)

const Bucket = "cloud.aws.s3.bucket"

func init() {
	config.Add(Bucket, "example", "s3 example bucket")
}

func main() {

	// parse config
	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// create background context
	ctx := context.Background()

	// start logrus
	logrus.Start()

	// get logrus instance from context
	l := logrus.FromContext(ctx)

	// create default aws config
	awsConfig := v2.NewDefaultConfig(ctx)

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
			log.Fatal(err)
		}

		log.Fatalf("unable check file %s in s3 bucket %s", filename, bucket)
	}

	l = l.WithField("lastModified", head.LastModified).
		WithField("versionId", head.VersionId)

	l.Debugf("file %s exists on bucket %s", filename, bucket)

}
