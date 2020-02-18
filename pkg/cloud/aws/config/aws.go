package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const (
	Key            = "cloud.aws.access.key.id"
	Secret         = "cloud.aws.secret.access.key"
	Region         = "cloud.aws.default.region"
	Session        = "cloud.aws.session.token"
	CustomEndpoint = "cloud.aws.custom.endpoint"
)

func init() {

	log.Println("getting configurations for aws")

	config.Add(Key, "", "defines the aws key")
	config.Add(Secret, "", "defines the aws secret")
	config.Add(Region, "", "defines the aws region")
	config.Add(Session, "", "defines the aws session token")
	config.Add(CustomEndpoint, false, "defines if should point to localhost")

}
