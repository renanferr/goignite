package elasticsearch

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
)

type ClientChecker struct {
	client *elasticsearch.Client
}

func (c *ClientChecker) Check(ctx context.Context) error {
	_, err := c.client.Ping(c.client.Ping.WithPretty())
	return err
}

func NewClientChecker(client *elasticsearch.Client) *ClientChecker {
	return &ClientChecker{client: client}
}
