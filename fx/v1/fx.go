package gifx

import (
	gilog "github.com/b2wdigital/goignite/v2/log"
	"go.uber.org/fx"
)

func New(opts ...fx.Option) *fx.App {
	opts = append(opts, fx.Logger(gilog.GetLogger()))
	return New(opts...)
}
