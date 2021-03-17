package newrelic

import (
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	"github.com/b2wdigital/goignite/v2/newrelic/v3"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				newrelic.NewApplication,
			),
		)
	})

	return options
}
