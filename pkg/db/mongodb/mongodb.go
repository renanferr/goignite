package mongodb

import (
	"context"
	"log"
	"strings"

	"github.com/jpfaria/goignite/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const Uri = "mongodb.uri"

func init() {

	log.Println("getting configurations for mongodb")

	config.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")

}

func NewClient( uri string ) (client *mongo.Client, database *mongo.Database, err error) {

	ctx := context.Background()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	connFields, err := connstring.Parse(uri)

	if err != nil {
		return nil, nil, err
	}

	database = client.Database(connFields.Database)

	log.Printf("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	return client, database, err
}

func NewDefaultClient() (*mongo.Client, *mongo.Database, error) {

	if err := config.Parse(); err != nil {
		return nil, nil, err
	}

	uri := config.Instance.String(Uri)

	return NewClient(uri)
}