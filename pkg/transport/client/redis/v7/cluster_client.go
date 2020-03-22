package redis

import (
	"context"
	"strings"

	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/go-redis/redis/v7"
)

func NewClusterClient(ctx context.Context, o *Options) (client *redis.ClusterClient, err error) {

	l := log.FromContext(ctx)

	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              o.Cluster.Addrs,
		MaxRedirects:       o.Cluster.MaxRedirects,
		ReadOnly:           o.Cluster.ReadOnly,
		RouteByLatency:     o.Cluster.RouteByLatency,
		RouteRandomly:      o.Cluster.RouteRandomly,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})

	ping := client.Ping()
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	l.Infof("Connected to Redis Cluster server: %s status: %s", strings.Join(client.Options().Addrs, ","), ping.String())

	if o.Health.Enabled {
		configureClusterHealthCheck(client, o)
	}

	return client, err
}

func NewDefaultClusterClient(ctx context.Context) (*redis.ClusterClient, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClusterClient(ctx, o)
}

func configureClusterHealthCheck(client *redis.ClusterClient, o *Options) {
	mc := NewClusterClientChecker(client)
	hc := health.NewHealthChecker("redis", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
