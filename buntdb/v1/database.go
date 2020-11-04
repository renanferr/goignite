package gibuntdb

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/tidwall/buntdb"
)

func NewDatabase(ctx context.Context, o *Options) (db *buntdb.DB, err error) {

	l := gilog.FromContext(ctx)

	var syncPolicy buntdb.SyncPolicy

	switch o.SyncPolicy {
	case 1:
		syncPolicy = 1
	case 2:
		syncPolicy = 2
	default:
		syncPolicy = 0
	}

	config := buntdb.Config{
		SyncPolicy:           syncPolicy,
		AutoShrinkPercentage: o.AutoShrink.Percentage,
		AutoShrinkMinSize:    o.AutoShrink.MinSize,
		AutoShrinkDisabled:   o.AutoShrink.Disabled,
	}

	db, err = buntdb.Open(o.Path)
	if err != nil {
		return nil, err
	}

	err = db.SetConfig(config)
	if err != nil {
		return nil, err
	}

	l.Infof("Connected to buntdb: %s", o.Path)

	return db, err

}

func NewDefaultDatabase(ctx context.Context) (*buntdb.DB, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewDatabase(ctx, o)
}
