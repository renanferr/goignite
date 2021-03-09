package gifreecache

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/coocood/freecache"
)

func NewCache(ctx context.Context, o *Options) (cache *freecache.Cache, err error) {

	logger := gilog.FromContext(ctx)

	cache = freecache.NewCache(o.CacheSize)

	logger.Infof("Created cache with size %v", o.CacheSize)

	return cache, err

}

func NewDefaultCache(ctx context.Context) (*freecache.Cache, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewCache(ctx, o)
}
