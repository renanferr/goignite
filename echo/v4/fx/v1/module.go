package giechofx

import (
	"context"
	"sync"

	gicontextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	giserver "github.com/b2wdigital/goignite/v2/server"
	giserverfx "github.com/b2wdigital/goignite/v2/server/fx/v1"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	Exts []giecho.Ext `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			gicontextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *giecho.Server {
					return giecho.NewDefault(ctx, p.Exts...)
				},
				func(srv *giecho.Server) *echo.Echo {
					return srv.Echo()
				},
			),
			fx.Provide(
				fx.Annotated{
					Group: giserverfx.ServersGroupKey,
					Target: func(srv *giecho.Server) giserver.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
