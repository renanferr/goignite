package mongodb

import (
	"context"
	"strings"

	"github.com/b2wdigital/goignite/pkg/config"
	h "github.com/b2wdigital/goignite/pkg/db/mongodb/health"
	"github.com/b2wdigital/goignite/pkg/db/mongodb/model"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func NewClient(ctx context.Context, o model.Options) (client *mongo.Client, database *mongo.Database, err error) {

	log := logrus.FromContext(ctx)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(o.Uri))

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

	log.Printf("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	if o.Health.Enabled {
		configureHealthCheck(client, o)
	}

	return client, database, err
}

func NewDefaultClient(ctx context.Context) (*mongo.Client, *mongo.Database, error) {

	log := logrus.FromContext(ctx)

	o := model.Options{}

	err := config.UnmarshalWithPath("db.mongodb", &o)
	if err != nil {
		log.Fatal(err)
	}

	return NewClient(ctx, o)
}

func configureHealthCheck(client *mongo.Client, o model.Options) {
	mc := h.NewMongoChecker(client)
	hc := health.NewHealthChecker("mongodb", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
