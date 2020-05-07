package godror

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	DataSourceName  = "transport.client.godror.datasourcename"
	ConnMaxLifetime = "transport.client.godror.connmaxlifetime"
	MaxIdleConns    = "transport.client.godror.maxidleconns"
	MaxOpenConns    = "transport.client.godror.maxopenconns"
)

func init() {

	log.Println("getting configurations for oracle (godror)")

	config.Add(DataSourceName, "", "database name and connection information")
	config.Add(ConnMaxLifetime, 0*time.Second, "sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever")
	config.Add(MaxIdleConns, 2, "sets the maximum number of connections in the idle connection pool.")
	config.Add(MaxOpenConns, 5, "sets the maximum number of open connections to the database.")

}
