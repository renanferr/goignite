package gifiberfx

import (
	"context"
	"sync"

	gicontextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
	giserver "github.com/b2wdigital/goignite/v2/server"
	giserverfx "github.com/b2wdigital/goignite/v2/server/fx/v1"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	Exts []gifiber.Ext `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			gicontextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *gifiber.Server {
					return gifiber.NewDefault(ctx, p.Exts...)
				},
				func(srv *gifiber.Server) *fiber.App {
					return srv.App()
				},
			),
			fx.Provide(
				fx.Annotated{
					Group: giserverfx.ServersGroupKey,
					Target: func(srv *gifiber.Server) giserver.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
