package fx

import (
	"context"
	"sync"

	"github.com/b2wdigital/goignite/v2/chi/v5"
	contextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	"github.com/b2wdigital/goignite/v2/server"
	serverfx "github.com/b2wdigital/goignite/v2/server/fx/v1"
	c "github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var once sync.Once

type params struct {
	fx.In
	Exts []chi.Ext `optional:"true"`
}

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *chi.Server {
					return chi.NewDefault(ctx, p.Exts...)
				},
				func(srv *chi.Server) *c.Mux {
					return srv.Mux()
				},
				fx.Annotated{
					Group: serverfx.ServersGroupKey,
					Target: func(srv *chi.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
