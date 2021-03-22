package server

import (
	"context"
	"sync"

	context2 "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/modules/context"
	server2 "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/modules/core/server"
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/server"
	s "github.com/b2wdigital/goignite/v2/core/server"
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
			context2.Module(),
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
					Group: server2.ServersGroupKey,
					Target: func(srv *server.Server) s.Server {
						return srv
					},
				},
			),
		)

	})

	return options
}
