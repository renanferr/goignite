package nats

import (
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	ginats "github.com/b2wdigital/goignite/v2/contrib/nats-io/nats.go.v1"
	"go.uber.org/fx"
)

var subsOnce sync.Once

func Module() fx.Option {
	options := fx.Options()

	subsOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				ginats.NewDefaultSubscriber,
			),
		)
	})

	return options
}
