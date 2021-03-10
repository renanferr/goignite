package gifx

import (
	gilog "github.com/b2wdigital/goignite/v2/log"
	"go.uber.org/fx"
)

func New(opts ...fx.Option) *fx.App {
	logger := gilog.GetLogger()
	opts = append([]fx.Option{fx.Logger(logger)}, opts...)
	return fx.New(opts...)
}
