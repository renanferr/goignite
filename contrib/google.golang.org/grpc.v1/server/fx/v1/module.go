package fx

import (
	"context"
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/context/fx/v1"
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/server"
	s "github.com/b2wdigital/goignite/v2/core/server"
	serverfx "github.com/b2wdigital/goignite/v2/core/server/fx/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type params struct {
	fx.In
	Exts []server.Ext `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *server.Server {
					return server.NewDefault(ctx, p.Exts...)
				},
				func(srv *server.Server) *grpc.Server {
					return srv.Server()
				},
				func(srv *server.Server) grpc.ServiceRegistrar {
					return srv.ServiceRegistrar()
				},
				fx.Annotated{
					Group: serverfx.ServersGroupKey,
					Target: func(srv *server.Server) s.Server {
						return srv
					},
				},
			),
		)

	})

	return options
}
