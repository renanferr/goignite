package giaws

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Key            = "aws.access.key.id"
	Secret         = "aws.secret.access.key"
	Region         = "aws.default.region"
	Session        = "aws.session.token"
	CustomEndpoint = "aws.custom.endpoint"
)

func init() {

	log.Println("getting configurations for aws")

	giconfig.Add(Key, "", "defines the aws key")
	giconfig.Add(Secret, "", "defines the aws secret")
	giconfig.Add(Region, "", "defines the aws region")
	giconfig.Add(Session, "", "defines the aws session token")
	giconfig.Add(CustomEndpoint, false, "defines if should point to localhost")

}
