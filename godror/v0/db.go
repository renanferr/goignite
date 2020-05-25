package gigodror

import (
	"context"
	"database/sql"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	_ "github.com/godror/godror"
)

const (
	TopicDB = "topic:godror:db"
)

func NewDB(ctx context.Context, o *Options) (db *sql.DB, err error) {

	l := gilog.FromContext(ctx)

	db, err = sql.Open("godror", o.DataSourceName)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	db.SetConnMaxLifetime(o.ConnMaxLifetime)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetMaxOpenConns(o.MaxOpenConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	gieventbus.Publish(TopicDB, db)

	l.Info("Connected to Oracle (godror) server")

	return db, err
}

func NewDefaultDB(ctx context.Context) (*sql.DB, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewDB(ctx, o)
}
