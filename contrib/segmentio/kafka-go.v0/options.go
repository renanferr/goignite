package kafka

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

type Options struct {
	Address   string
	Topic     string
	Partition int
	Network   string
}

func CustomOptions(path string) (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(path, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func DefaultOptions() (*Options, error) {
	return CustomOptions(root)
}
