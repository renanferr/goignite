package mongodb

import (
	"context"
	"strings"

	"github.com/b2wdigital/goignite/pkg/errors"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func NewClientWithIntegrations(ctx context.Context, o *Options, integrations []Integrator) (client *mongo.Client, database *mongo.Database, err error) {

	co := clientOptions(o)

	for _, integrator := range integrations {
		err = integrator.Integrate(ctx, co)
		if err != nil {
			return nil, nil, errors.Wrap(err, errors.Internalf("error on integrate mongodb"))
		}
	}

	return newClient(ctx, co, o)
}

func NewClient(ctx context.Context, o *Options) (client *mongo.Client, database *mongo.Database, err error) {

	co := clientOptions(o)

	return newClient(ctx, co, o)
}

func NewDefaultClientWithIntegrations(ctx context.Context, integrations []Integrator) (*mongo.Client, *mongo.Database, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClientWithIntegrations(ctx, o, integrations)
}

func NewDefaultClient(ctx context.Context) (*mongo.Client, *mongo.Database, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}

func newClient(ctx context.Context, co *options.ClientOptions, o *Options) (client *mongo.Client, database *mongo.Database, err error) {

	l := log.FromContext(ctx)

	client, err = mongo.Connect(ctx, co)

	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	connFields, err := connstring.Parse(o.Uri)

	if err != nil {
		return nil, nil, err
	}

	database = client.Database(connFields.Database)

	l.Infof("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	if o.Health.Enabled {
		configureHealthCheck(client, o)
	}

	return client, database, err
}

func clientOptions(o *Options) *options.ClientOptions {
	return options.Client().ApplyURI(o.Uri)
}

func configureHealthCheck(client *mongo.Client, o *Options) {
	mc := NewClientChecker(client)
	hc := health.NewHealthChecker("mongodb", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
