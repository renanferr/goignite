package giserverfx

import (
	"context"
	"sync"

	gicobra "github.com/b2wdigital/goignite/v2/cobra/v1"
	gicontextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	giserver "github.com/b2wdigital/goignite/v2/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const (
	ServersGroupKey = "_gi_server_servers_"
)

type srvParams struct {
	fx.In
	Servers []giserver.Server `group:"_gi_server_servers_"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			gicontextfx.Module(),
			fx.Invoke(
				func(ctx context.Context, p srvParams) error {

					return gicobra.Run(
						&cobra.Command{
							Run: func(cmd *cobra.Command, args []string) {
								giserver.Serve(ctx, p.Servers...)
							},
						},
					)

				},
			),
		)
	})

	return options
}
