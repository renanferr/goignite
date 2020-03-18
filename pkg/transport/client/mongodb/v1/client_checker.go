package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ClientChecker struct {
	client *mongo.Client
}

func (c *ClientChecker) Check(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

func NewClientChecker(client *mongo.Client) *ClientChecker {
	return &ClientChecker{client: client}
}
