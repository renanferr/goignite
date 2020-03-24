package memdb

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/hashicorp/go-memdb"
)

func NewDatabase(ctx context.Context, schema *memdb.DBSchema) (db *memdb.MemDB, err error) {

	l := log.FromContext(ctx)

	db, err = memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	l.Info("Connected to go-memdb")

	return db, err

}
