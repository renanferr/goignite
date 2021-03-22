package fx

import (
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/context/fx/v1"
	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				datadog.NewTracer,
			),
		)
	})

	return options
}
