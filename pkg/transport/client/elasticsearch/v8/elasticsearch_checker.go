package elasticsearch

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearchChecker struct {
	client *elasticsearch.Client
}

func (c *ElasticSearchChecker) Check(ctx context.Context) error {
	_, err := c.client.Ping(c.client.Ping.WithPretty())
	return err
}

func NewElasticSearchChecker(client *elasticsearch.Client) *ElasticSearchChecker {
	return &ElasticSearchChecker{client: client}
}
