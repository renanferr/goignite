package giredis

import (
	"context"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-redis/redis/v7"
)

const (
	TopicClient = "topic:redis:client"
)

func NewClient(ctx context.Context, o *Options) (client *redis.Client, err error) {

	l := gilog.FromContext(ctx)

	client = redis.NewClient(&redis.Options{
		Network:            o.Client.Network,
		Addr:               o.Client.Addr,
		DB:                 o.Client.DB,
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

	ping := client.Conn().Ping()
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	gieventbus.Publish(TopicClient, client)

	l.Infof("Connected to Redis server: %s %s", client.Options().Addr, ping.String())

	return client, err
}

func NewDefaultClient(ctx context.Context) (*redis.Client, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}