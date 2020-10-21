package gielasticsearch

import (
	"context"
	"strings"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

const (
	TopicClient = "topic:gielasticsearch:client"
)

func NewClient(ctx context.Context, o *Options) (client *elasticsearch.Client, err error) {

	l := gilog.FromContext(ctx)

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

	gieventbus.Publish(TopicClient, client)

	l.Infof("Connected to Elastic Search server: %v status: %s", strings.Join(o.Addresses, ","), res.Status())

	return client, err
}

func backOff(attempt int) time.Duration {
	b := giconfig.Duration(RetryBackoff)
	return time.Duration(attempt) * b
}

func NewDefaultClient(ctx context.Context) (*elasticsearch.Client, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}
