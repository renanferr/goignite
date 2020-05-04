package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Integrator interface {
	Integrate(context.Context, *options.ClientOptions) error
}
