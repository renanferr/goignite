package ginewrelicfx

import (
	"sync"

	gicontext "github.com/b2wdigital/goignite/v2/context/fx/v1"
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			gicontext.Module(),
			fx.Invoke(
				ginewrelic.NewApplication,
			),
		)
	})

	return options
}
