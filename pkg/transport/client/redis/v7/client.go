package redis

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/errors"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/go-redis/redis/v7"
)

func NewClientWithIntegrations(ctx context.Context, o *Options, integrations []Integrator) (client *redis.Client, err error) {

	client, err = NewClient(ctx, o)
	if err != nil {
		return client, err
	}

	for _, integrator := range integrations {
		err = integrator.Integrate(ctx, client)
		if err != nil {
			return nil, errors.Wrap(err, errors.Internalf("error on integrate mongodb"))
		}
	}

	return client, err
}

func NewClient(ctx context.Context, o *Options) (client *redis.Client, err error) {

	l := log.FromContext(ctx)

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

	l.Infof("Connected to Redis server: %s %s", client.Options().Addr, ping.String())

	if o.Health.Enabled {
		configureHealthCheck(client, o)
	}

	return client, err
}

func NewDefaultClientWithIntegrations(ctx context.Context, integrations []Integrator) (*redis.Client, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClientWithIntegrations(ctx, o, integrations)
}

func NewDefaultClient(ctx context.Context) (*redis.Client, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}

func configureHealthCheck(client *redis.Client, o *Options) {
	mc := NewClientChecker(client)
	hc := health.NewHealthChecker("redis", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
