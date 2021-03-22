package elasticsearch

import (
	"context"
	"strings"
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Ext func(context.Context, *elasticsearch.Client) error

func NewClient(ctx context.Context, o *Options, exts ...Ext) (client *elasticsearch.Client, err error) {

	logger := log.FromContext(ctx)

	cfg := elasticsearch.Config{
		Addresses:             o.Addresses,
		Username:              o.Username,
		Password:              o.Password,
		CloudID:               o.CloudID,
		APIKey:                o.APIKey,
		RetryOnStatus:         o.RetryOnStatus,
		DisableRetry:          o.DisableRetry,
		EnableRetryOnTimeout:  o.EnableRetryOnTimeout,
		MaxRetries:            o.MaxRetries,
		DiscoverNodesOnStart:  o.DiscoverNodesOnStart,
		DiscoverNodesInterval: o.DiscoverNodesInterval,
		EnableMetrics:         o.EnableMetrics,
		EnableDebugLogger:     o.EnableDebugLogger,
		RetryBackoff:          backOff,
		Logger:                &Logger{},
	}

	if o.CACert != "" {
		cfg.CACert = []byte(o.CACert)
	}

	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	var res *esapi.Response

	res, err = client.Ping(client.Ping.WithPretty())
	if err != nil {
		return nil, err
	}

	for _, ext := range exts {
		if err := ext(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to Elastic Search server: %v status: %s", strings.Join(o.Addresses, ","), res.Status())

	return client, err
}

func backOff(attempt int) time.Duration {
	b := config.Duration(retryBackoff)
	return time.Duration(attempt) * b
}

func NewDefaultClient(ctx context.Context, exts ...Ext) (*elasticsearch.Client, error) {

	logger := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClient(ctx, o, exts...)
}
