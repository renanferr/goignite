package giredis

import (
	"context"
	"strings"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-redis/redis/v8"
)

const (
	TopicClusterClient = "topic:redis:v8:cluster:client"
)

func NewClusterClient(ctx context.Context, o *Options) (client *redis.ClusterClient, err error) {

	logger := gilog.FromContext(ctx)

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

	ping := client.Ping(ctx)
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	gieventbus.Publish(TopicClusterClient, client)

	logger.Infof("Connected to Redis Cluster server: %s status: %s", strings.Join(client.Options().Addrs, ","), ping.String())

	return client, err
}

func NewDefaultClusterClient(ctx context.Context) (*redis.ClusterClient, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClusterClient(ctx, o)
}
