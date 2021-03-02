package gigodror

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root            = "gi.godror"
	dataSourceName  = root + ".dataSourceName"
	connMaxLifetime = root + ".connMaxLifetime"
	maxIdleConns    = root + ".maxIdleConns"
	maxOpenConns    = root + ".maxOpenConns"
	ExtRoot         = root + ".Ext"
)

func init() {

	giconfig.Add(dataSourceName, "", "database name and connection information")
	giconfig.Add(connMaxLifetime, 0*time.Second, "sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever")
	giconfig.Add(maxIdleConns, 2, "sets the maximum number of connections in the idle connection pool.")
	giconfig.Add(maxOpenConns, 5, "sets the maximum number of open connections to the database.")
}
