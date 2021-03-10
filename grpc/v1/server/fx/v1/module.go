package gigrpcfx

import (
	"context"
	"sync"

	gicontext "github.com/b2wdigital/goignite/v2/context/fx/v1"
	gigrpc "github.com/b2wdigital/goignite/v2/grpc/v1/server"
	giserver "github.com/b2wdigital/goignite/v2/server"
	giserverfx "github.com/b2wdigital/goignite/v2/server/fx/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type params struct {
	fx.In
	Exts []gigrpc.Ext `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			gicontext.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *gigrpc.Server {
					return gigrpc.NewDefault(ctx, p.Exts...)
				},
				func(srv *gigrpc.Server) *grpc.Server {
					return srv.Server()
				},
				func(srv *gigrpc.Server) grpc.ServiceRegistrar {
					return srv.ServiceRegistrar()
				},
				fx.Annotated{
					Group: giserverfx.ServersGroupKey,
					Target: func(srv *gigrpc.Server) giserver.Server {
						return srv
					},
				},
			),
		)

	})

	return options
}
