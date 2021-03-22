package fiber

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/gofiber/fiber/v2"
)

type Options struct {
	Port   int
	Config *fiber.Config
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
