package fx

import (
	"context"
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	"github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/b2wdigital/goignite/v2/server"
	serverfx "github.com/b2wdigital/goignite/v2/server/fx/v1"
	e "github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	Exts []echo.Ext `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *echo.Server {
					return echo.NewDefault(ctx, p.Exts...)
				},
				func(srv *echo.Server) *e.Echo {
					return srv.Echo()
				},
			),
			fx.Provide(
				fx.Annotated{
					Group: serverfx.ServersGroupKey,
					Target: func(srv *echo.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
