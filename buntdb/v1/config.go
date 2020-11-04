package gibuntdb

import (
	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	Path                 = "gi.buntdb.path"
	SyncPolicy           = "gi.buntdb.syncPolicy"
	AutoShrinkPercentage = "gi.buntdb.autoShrink.percentage"
	AutoShrinkMinSize    = "gi.buntdb.autoShrink.minSize"
	AutoShrinkDisabled   = "gi.buntdb.autoShrink.disabled"
)

func init() {

	log.Println("getting configurations for buntdb")

	giconfig.Add(Path, ":memory:", "open opens a database at the provided path")
	giconfig.Add(SyncPolicy, 1, "adjusts how often the data is synced to disk (Never: 0, EverySecond: 1, Always: 2)")
	giconfig.Add(AutoShrinkPercentage, 100, "is used by the background process to trigger a shrink of the aof file when the size of the file is larger than the percentage of the result of the previous shrunk file")
	giconfig.Add(AutoShrinkMinSize, 32*1024*102, "defines the minimum size of the aof file before an automatic shrink can occur")
	giconfig.Add(AutoShrinkDisabled, false, "turns off automatic background shrinking")

}
