package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Integrator interface {
	Integrate(context.Context, *aws.Config) error
}
