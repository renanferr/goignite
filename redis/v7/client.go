package giredis

import (
	"context"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-redis/redis/v7"
)

const (
	TopicClient = "topic:redis:v7:client"
)

func NewClient(ctx context.Context, o *Options) (client *redis.Client, err error) {

	l := gilog.FromContext(ctx)

	if redisSentinel(o) {
		client = failOverClient(o)
	} else {
		client = standaloneClient(o)
	}

	ping := client.Conn().Ping()
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	gieventbus.Publish(TopicClient, client)

	l.Infof("Connected to Redis server: %s %s", client.Options().Addr, ping.String())

	return client, err
}

func failOverClient(o *Options) *redis.Client {
	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:         o.Sentinel.MasterName,
		SentinelAddrs:      o.Sentinel.Addrs,
		SentinelPassword:   o.Sentinel.Password,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		DB:                 o.Client.DB,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})
}

func standaloneClient(o *Options) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               o.Client.Addr,
		Network:            o.Client.Network,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		DB:                 o.Client.DB,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})
}

func redisSentinel(o *Options) bool {
	return o.Sentinel.MasterName != "" || o.Sentinel.Addrs != nil
}

func NewDefaultClient(ctx context.Context) (*redis.Client, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}
