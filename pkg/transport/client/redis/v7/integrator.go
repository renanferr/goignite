package redis

import (
	"context"

	"github.com/go-redis/redis/v7"
)

type Integrator interface {
	Integrate(context.Context, *redis.Client) error
}
