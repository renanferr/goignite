package buntdb

import (
	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Path                 = "transport.client.buntdb.parh"
	SyncPolicy           = "transport.client.buntdb.syncpolicy"
	AutoShrinkPercentage = "transport.client.buntdb.autoshrink.percentage"
	AutoShrinkMinSize    = "transport.client.buntdb.autoshrink.minsize"
	AutoShrinkDisabled   = "transport.client.buntdb.autoshrink.disabled"
)

func init() {

	log.Println("getting configurations for buntdb")

	config.Add(Path, ":memory:", "open opens a database at the provided path")
	config.Add(SyncPolicy, 1, "adjusts how often the data is synced to disk (Never: 0, EverySecond: 1, Always: 2)")
	config.Add(AutoShrinkPercentage, 100, "is used by the background process to trigger a shrink of the aof file when the size of the file is larger than the percentage of the result of the previous shrunk file")
	config.Add(AutoShrinkMinSize, 32*1024*102, "defines the minimum size of the aof file before an automatic shrink can occur")
	config.Add(AutoShrinkDisabled, false, "turns off automatic background shrinking")

}
