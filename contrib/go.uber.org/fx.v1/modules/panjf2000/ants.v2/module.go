package ants

import (
	"sync"

	ants "github.com/b2wdigital/goignite/v2/contrib/panjf2000/ants.v2"
	a "github.com/panjf2000/ants/v2"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			fx.Provide(
				func(pool *a.Pool, m []ants.Middleware) *ants.Wrapper {
					return ants.New(pool, m...)
				},
			),
		)
	})

	return options
}
