package gigodror

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	DataSourceName  = "gi.godror.dataSourceName"
	ConnMaxLifetime = "gi.godror.connMaxLifetime"
	MaxIdleConns    = "gi.godror.maxIdleConns"
	MaxOpenConns    = "gi.godror.maxOpenConns"
)

func init() {

	giconfig.Add(DataSourceName, "", "database name and connection information")
	giconfig.Add(ConnMaxLifetime, 0*time.Second, "sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever")
	giconfig.Add(MaxIdleConns, 2, "sets the maximum number of connections in the idle connection pool.")
	giconfig.Add(MaxOpenConns, 5, "sets the maximum number of open connections to the database.")
}
