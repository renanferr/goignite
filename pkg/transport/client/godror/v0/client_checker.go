package godror

import (
	"context"
	"database/sql"
)

type ClientChecker struct {
	db *sql.DB
}

func (c *ClientChecker) Check(ctx context.Context) error {
	if err := c.db.Ping(); err != nil {
		return err
	}
	return nil
}

func NewClientChecker(db *sql.DB) *ClientChecker {
	return &ClientChecker{db: db}
}
