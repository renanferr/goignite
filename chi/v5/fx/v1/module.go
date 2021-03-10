package gichifx

import (
	"context"
	"sync"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gicontextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	giserver "github.com/b2wdigital/goignite/v2/server"
	giserverfx "github.com/b2wdigital/goignite/v2/server/fx/v1"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var once sync.Once

type params struct {
	fx.In
	Exts []gichi.Ext `optional:"true"`
}

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			gicontextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *gichi.Server {
					return gichi.NewDefault(ctx, p.Exts...)
				},
				func(srv *gichi.Server) *chi.Mux {
					return srv.Mux()
				},
				fx.Annotated{
					Group: giserverfx.ServersGroupKey,
					Target: func(srv *gichi.Server) giserver.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
