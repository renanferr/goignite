package gocql

import (
	"context"

	"github.com/gocql/gocql"
)

type SessionChecker struct {
	session *gocql.Session
}

func (c *SessionChecker) Check(ctx context.Context) error {
	return c.session.Query("void").Exec()
}

func NewSessionChecker(session *gocql.Session) *SessionChecker {
	return &SessionChecker{session: session}
}
