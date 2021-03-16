package gifx

import (
	"go.uber.org/fx"
)

func New(opts ...fx.Option) *fx.App {
	opts = append([]fx.Option{fx.Logger(NewLogger())}, opts...)
	return fx.New(opts...)
}
