package gimongo

import (
	"context"
	"strings"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"go.mongodb.org/mongo-driver/event"
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

	co := clientOptions(ctx, o)

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

	logger := gilog.FromContext(ctx)

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

	logger.Infof("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	return client, database, err
}

func clientOptions(ctx context.Context, o *Options) *options.ClientOptions {

	logger := gilog.FromContext(ctx)

	clientOptions := options.Client().ApplyURI(o.Uri)
	clientOptions.SetMonitor(&event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			logger.Debugf("mongodb cmd - %v %s %s %v", startedEvent.ConnectionID, startedEvent.CommandName, startedEvent.DatabaseName, startedEvent.RequestID)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			logger.Debugf("mongodb cmd - %v %s %vus %v", succeededEvent.ConnectionID, succeededEvent.CommandName, succeededEvent.DurationNanos, succeededEvent.RequestID)
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			logger.Errorf("mongodb cmd - %v %s %s %v", failedEvent.ConnectionID, failedEvent.CommandName, failedEvent.Failure, failedEvent.RequestID)
		},
	})
	clientOptions.SetPoolMonitor(&event.PoolMonitor{
		Event: func(poolEvent *event.PoolEvent) {
			logger.Debugf("mongodb conn pool - %v %s %s %s", poolEvent.ConnectionID, poolEvent.Type, poolEvent.Reason, poolEvent.Address)
		},
	})

	if o.Auth != nil {
		setAuthOptions(o, clientOptions)
	}

	return clientOptions
}

func setAuthOptions(o *Options, clientOptions *options.ClientOptions) {

	if o.Auth.Password == "" && o.Auth.Username == "" {
		return
	}

	if clientOptions.Auth == nil {
		clientOptions.Auth = &options.Credential{}
	}

	if o.Auth.Password != "" {
		clientOptions.Auth.Password = o.Auth.Password
		clientOptions.Auth.PasswordSet = true
	}

	if o.Auth.Username != "" {
		clientOptions.Auth.Username = o.Auth.Username
	}

	if clientOptions.Auth.AuthSource == "" {
		connFields, _ := connstring.Parse(clientOptions.GetURI())
		clientOptions.Auth.AuthSource = connFields.Database
	}
}
