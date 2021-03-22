package freecache

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/coocood/freecache"
)

func NewCache(ctx context.Context, o *Options) (cache *freecache.Cache, err error) {

	logger := log.FromContext(ctx)

	cache = freecache.NewCache(o.CacheSize)

	logger.Infof("Created cache with size %v", o.CacheSize)

	return cache, err

}

func NewDefaultCache(ctx context.Context, opts ...Option) (*freecache.Cache, error) {

	logger := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	for _, opt := range opts {
		opt(o)
	}

	return NewCache(ctx, o)
}
