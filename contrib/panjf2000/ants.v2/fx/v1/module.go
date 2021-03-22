package fx

import (
	"sync"

	giants "github.com/b2wdigital/goignite/v2/contrib/panjf2000/ants.v2"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			fx.Provide(
				func(pool *ants.Pool, m []giants.Middleware) *giants.Wrapper {
					return giants.New(pool, m...)
				},
			),
		)
	})

	return options
}
