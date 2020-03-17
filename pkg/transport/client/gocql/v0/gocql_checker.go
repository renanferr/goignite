package gocql

import (
	"context"

	"github.com/gocql/gocql"
)

type CassandraChecker struct {
	session *gocql.Session
}

func (c *CassandraChecker) Check(ctx context.Context) error {
	return c.session.Query("void").Exec()
}

func NewCassandraChecker(session *gocql.Session) *CassandraChecker {
	return &CassandraChecker{session: session}
}
