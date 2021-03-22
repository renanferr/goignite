package chi

import (
	"context"
	"sync"

	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	context2 "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/modules/context"
	server2 "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/modules/core/server"
	"github.com/b2wdigital/goignite/v2/core/server"
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
			context2.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *chi.Server {
					return chi.NewDefault(ctx, p.Exts...)
				},
				func(srv *chi.Server) *c.Mux {
					return srv.Mux()
				},
				fx.Annotated{
					Group: server2.ServersGroupKey,
					Target: func(srv *chi.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
