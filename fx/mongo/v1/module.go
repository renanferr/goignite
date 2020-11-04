package gifxmongo

import (
	"sync"

	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				gimongo.DefaultOptions,
				gimongo.NewClient,
			),
		)
	})

	return options
}
