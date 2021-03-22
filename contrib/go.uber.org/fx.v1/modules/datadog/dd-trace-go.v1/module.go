package datadog

import (
	"sync"

	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	"github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/modules/context"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			context.Module(),
			fx.Invoke(
				datadog.NewTracer,
			),
		)
	})

	return options
}
