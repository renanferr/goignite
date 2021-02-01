package gigocql

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Hosts                    = "gi.gocql.hosts"
	Port                     = "gi.gocql.port"
	Username                 = "gi.gocql.username"
	Password                 = "gi.gocql.password"
	CQLVersion               = "gi.gocql.CQLVersion"
	ProtoVersion             = "gi.gocql.protoVersion"
	Timeout                  = "gi.gocql.timeout"
	ConnectTimeout           = "gi.gocql.connecttimeout"
	Keyspace                 = "gi.gocql.keyspace"
	NumConns                 = "gi.gocql.numConns"
	Consistency              = "gi.gocql.consistency"
	SocketKeepalive          = "gi.gocql.socketKeepAlive"
	MaxPreparedStmts         = "gi.gocql.maxPreparedStmts"
	MaxRoutingKeyInfo        = "gi.gocql.maxRoutingKeyInfo"
	PageSize                 = "gi.gocql.pageSize"
	DefaultTimestamp         = "gi.gocql.defaultTimestamp"
	ReconnectInterval        = "gi.gocql.reconnectInterval"
	MaxWaitSchemaAgreement   = "gi.gocql.maxWaitSchemaAgreement"
	DisableInitialHostLookup = "gi.gocql.disableInitialHostLookup"
	WriteCoalesceWaitTime    = "gi.gocql.writeCoalesceWaitTime"
)

func init() {

	giconfig.Add(Hosts, []string{"127.0.0.1"}, "addresses for the initial connections")
	giconfig.Add(Port, 9042, "define port")
	giconfig.Add(Username, "", "define username")
	giconfig.Add(Password, "", "define password")
	giconfig.Add(CQLVersion, "3.0.0", "define cql version")
	giconfig.Add(ProtoVersion, 0, "define version of the native protocol to use")
	giconfig.Add(Timeout, 600*time.Millisecond, "connection timeout")
	giconfig.Add(ConnectTimeout, 600*time.Millisecond, "initial connection timeout, used during initial dial to server")
	giconfig.Add(Keyspace, "", "initial keyspace (optional)")
	giconfig.Add(NumConns, 2, "number of connections per host")
	giconfig.Add(Consistency, "QUORUM", "default consistency level (default: Quorum) (values: ANY, ONE, TWO, THREE, QUORUM, ALL, LOCAL_QUORUM, EACH_QUORUM, LOCAL_ONE)")
	giconfig.Add(SocketKeepalive, 0*time.Millisecond, "The keepalive period to use, enabled if > 0 (default: 0)")
	giconfig.Add(MaxPreparedStmts, 1000, "Sets the maximum cache size for prepared statements globally for gocql")
	giconfig.Add(MaxRoutingKeyInfo, 1000, "Sets the maximum cache size for query info about statements for each session")
	giconfig.Add(PageSize, 5000, "Default page size to use for created sessions")
	giconfig.Add(DefaultTimestamp, true, "Sends a client side timestamp for all requests which overrides the timestamp at which it arrives at the server. (default: true, only enabled for protocol 3 and above)")
	giconfig.Add(ReconnectInterval, 10*time.Millisecond, "If not zero, gocql attempt to reconnect known DOWN nodes in every ReconnectInterval")
	giconfig.Add(MaxWaitSchemaAgreement, 60*time.Second, "The maximum amount of time to wait for schema agreement in a cluster after receiving a schema change frame")
	giconfig.Add(DisableInitialHostLookup, true, "If true then the driver will not attempt to get host info from the system.peers table")
	giconfig.Add(WriteCoalesceWaitTime, 200*time.Microsecond, "The time to wait for frames before flushing the frames connection to Cassandra")
}
