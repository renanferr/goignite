package freecache

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

type Options struct {
	CacheSize int
}

type Option func(options *Options)

func WithCacheSize(cacheSize int) Option {
	return func(options *Options) {
		options.CacheSize = cacheSize
	}
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

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
