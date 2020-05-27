package gimongo

import (
	"context"
	"strings"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	TopicClientOptions = "topic:gimongo:clientoptions"
	TopicClient        = "topic:gimongo:client"
	TopicDatabase      = "topic:gimongo:database"
)

func NewClient(ctx context.Context, o *Options) (client *mongo.Client, database *mongo.Database, err error) {

	co := clientOptions(o)

	gieventbus.Publish(TopicClientOptions, co)

	client, database, err = newClient(ctx, co)
	if err != nil {
		return nil, nil, err
	}

	gieventbus.Publish(TopicClient, client)
	gieventbus.Publish(TopicDatabase, database)

	return client, database, err
}

func NewDefaultClient(ctx context.Context) (*mongo.Client, *mongo.Database, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}

func newClient(ctx context.Context, co *options.ClientOptions) (client *mongo.Client, database *mongo.Database, err error) {

	l := gilog.FromContext(ctx)

	client, err = mongo.Connect(ctx, co)

	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	connFields, err := connstring.Parse(co.GetURI())

	if err != nil {
		return nil, nil, err
	}

	database = client.Database(connFields.Database)

	l.Infof("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	return client, database, err
}

func clientOptions(o *Options) *options.ClientOptions {
	return options.Client().ApplyURI(o.Uri)
}
