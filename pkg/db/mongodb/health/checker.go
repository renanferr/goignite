package health

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoChecker struct {
	client *mongo.Client
}

func (c *MongoChecker) Check(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

func NewMongoChecker(client *mongo.Client) *MongoChecker {
	return &MongoChecker{client: client}
}
