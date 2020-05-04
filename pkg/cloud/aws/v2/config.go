package aws

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Key             = "aws.access.key.id"
	Secret          = "aws.secret.access.key"
	Region          = "aws.default.region"
	Session         = "aws.session.token"
	CustomEndpoint  = "aws.custom.endpoint"
)

func init() {

	log.Println("getting configurations for aws")

	config.Add(Key, "", "defines the aws key")
	config.Add(Secret, "", "defines the aws secret")
	config.Add(Region, "", "defines the aws region")
	config.Add(Session, "", "defines the aws session token")
	config.Add(CustomEndpoint, false, "defines if should point to localhost")

}
