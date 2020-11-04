package gimemdb

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/hashicorp/go-memdb"
)

func NewDatabase(ctx context.Context, schema *memdb.DBSchema) (db *memdb.MemDB, err error) {

	l := gilog.FromContext(ctx)

	db, err = memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	l.Info("Connected to go-memdb")

	return db, err

}
