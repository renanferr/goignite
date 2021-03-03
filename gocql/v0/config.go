package gigocql

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root                     = "gi.gocql"
	hosts                    = root + ".hosts"
	port                     = root + ".port"
	dc                       = root + ".dc"
	username                 = root + ".username"
	password                 = root + ".password"
	cqlVersion               = root + ".CQLVersion"
	protoVersion             = root + ".protoVersion"
	timeout                  = root + ".timeout"
	connectTimeout           = root + ".connecttimeout"
	keyspace                 = root + ".keyspace"
	numConns                 = root + ".numConns"
	consistency              = root + ".consistency"
	socketKeepalive          = root + ".socketKeepAlive"
	maxPreparedStmts         = root + ".maxPreparedStmts"
	maxRoutingKeyInfo        = root + ".maxRoutingKeyInfo"
	pageSize                 = root + ".pageSize"
	defaultTimestamp         = root + ".defaultTimestamp"
	reconnectInterval        = root + ".reconnectInterval"
	maxWaitSchemaAgreement   = root + ".maxWaitSchemaAgreement"
	disableInitialHostLookup = root + ".disableInitialHostLookup"
	writeCoalesceWaitTime    = root + ".writeCoalesceWaitTime"
	ExtRoot                  = root + ".ext"
)

func init() {

	giconfig.Add(hosts, []string{"127.0.0.1"}, "addresses for the initial connections")
	giconfig.Add(port, 9042, "define port")
	giconfig.Add(dc, "", "define DC")
	giconfig.Add(username, "", "define username")
	giconfig.Add(password, "", "define password")
	giconfig.Add(cqlVersion, "3.0.0", "define cql version")
	giconfig.Add(protoVersion, 0, "define version of the native protocol to use")
	giconfig.Add(timeout, 600*time.Millisecond, "connection timeout")
	giconfig.Add(connectTimeout, 600*time.Millisecond, "initial connection timeout, used during initial dial to server")
	giconfig.Add(keyspace, "", "initial keyspace (optional)")
	giconfig.Add(numConns, 2, "number of connections per host")
	giconfig.Add(consistency, "QUORUM", "default consistency level (default: Quorum) (values: ANY, ONE, TWO, THREE, QUORUM, ALL, LOCAL_QUORUM, EACH_QUORUM, LOCAL_ONE)")
	giconfig.Add(socketKeepalive, 0*time.Millisecond, "The keepalive period to use, enabled if > 0 (default: 0)")
	giconfig.Add(maxPreparedStmts, 1000, "Sets the maximum cache size for prepared statements globally for gocql")
	giconfig.Add(maxRoutingKeyInfo, 1000, "Sets the maximum cache size for query info about statements for each session")
	giconfig.Add(pageSize, 5000, "Default page size to use for created sessions")
	giconfig.Add(defaultTimestamp, true, "Sends a client side timestamp for all requests which overrides the timestamp at which it arrives at the server. (default: true, only enabled for protocol 3 and above)")
	giconfig.Add(reconnectInterval, 10*time.Millisecond, "If not zero, gocql attempt to reconnect known DOWN nodes in every ReconnectInterval")
	giconfig.Add(maxWaitSchemaAgreement, 60*time.Second, "The maximum amount of time to wait for schema agreement in a cluster after receiving a schema change frame")
	giconfig.Add(disableInitialHostLookup, true, "If true then the driver will not attempt to get host info from the system.peers table")
	giconfig.Add(writeCoalesceWaitTime, 200*time.Microsecond, "The time to wait for frames before flushing the frames connection to Cassandra")
}
