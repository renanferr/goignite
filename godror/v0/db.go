package gigodror

import (
	"context"
	"database/sql"

	gilog "github.com/b2wdigital/goignite/log"
	_ "github.com/godror/godror"
)

type ext func(context.Context, *sql.DB) error

func NewDB(ctx context.Context, o *Options, exts ...ext) (db *sql.DB, err error) {

	logger := gilog.FromContext(ctx)

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

	for _, ext := range exts {
		if err := ext(ctx, db); err != nil {
			panic(err)
		}
	}

	logger.Info("Connected to Oracle (godror) server")

	return db, err
}

func NewDefaultDB(ctx context.Context, exts ...ext) (*sql.DB, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewDB(ctx, o, exts...)
}
