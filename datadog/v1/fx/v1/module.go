package gidatadogfx

import (
	"sync"

	gicontextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	gidatadog "github.com/b2wdigital/goignite/v2/datadog/v1"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			gicontextfx.Module(),
			fx.Invoke(
				gidatadog.NewTracer,
			),
		)
	})

	return options
}
