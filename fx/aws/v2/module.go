package gifxaws

import (
	"sync"

	giaws "github.com/b2wdigital/goignite/aws/v2"
	"go.uber.org/fx"
)

var once sync.Once

func AWSModule() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				giaws.NewDefaultConfig,
			),
		)
	})

	return options
}
