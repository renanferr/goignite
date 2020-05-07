package elasticsearch

import (
	"context"
	"database/sql"

	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	_ "github.com/godror/godror"
)

func NewDB(ctx context.Context, o *Options) (db *sql.DB, err error) {

	l := log.FromContext(ctx)

	db, err = sql.Open("godror", o.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.SetConnMaxLifetime(o.ConnMaxLifetime)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetMaxOpenConns(o.MaxOpenConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	l.Infof("Connected to Oracle (godror) server: %v")

	if o.Health.Enabled {
		configureHealthCheck(db, o)
	}

	return db, err
}

func NewDefaultDB(ctx context.Context) (*sql.DB, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewDB(ctx, o)
}

func configureHealthCheck(client *sql.DB, o *Options) {
	mc := NewClientChecker(client)
	hc := health.NewHealthChecker("oracle", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
