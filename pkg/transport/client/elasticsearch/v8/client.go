package elasticsearch

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func NewClient(ctx context.Context, o *Options) (client *elasticsearch.Client, err error) {

	l := log.FromContext(ctx)

	cfg := elasticsearch.Config{
		Addresses:             o.Addresses,
		Username:              o.Username,
		Password:              o.Password,
		CloudID:               o.CloudID,
		APIKey:                o.APIKey,
		RetryOnStatus:         stringToIntSlice(o.RetryOnStatus),
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

	l.Infof("Connected to Elastic Search server: %v status: %s", strings.Join(o.Addresses, ","), res.Status())

	if o.Health.Enabled {
		configureHealthCheck(client, o)
	}

	return client, err
}

func backOff(attempt int) time.Duration {
	b := config.Duration(RetryBackoff)
	return time.Duration(attempt) * b
}

// BUG: https://github.com/knadh/koanf/issues/24
func stringToIntSlice(values []string) []int {
	var valuesInt []int

	for i := range values {
		text := values[i]
		number, err := strconv.Atoi(text)
		if err != nil {
			log.Error(err)
			continue
		}
		valuesInt = append(valuesInt, number)
	}

	return valuesInt
}

func NewDefaultClient(ctx context.Context) (*elasticsearch.Client, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}

func configureHealthCheck(client *elasticsearch.Client, o *Options) {
	mc := NewClientChecker(client)
	hc := health.NewHealthChecker("elasticsearch", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
