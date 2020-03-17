package gocql

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Hosts                    = "transport.client.gocql.hosts"
	Port                     = "transport.client.gocql."
	Username                 = "transport.client.gocql.username"
	Password                 = "transport.client.gocql.password"
	CQLVersion               = "transport.client.gocql.cqlversion"
	ProtoVersion             = "transport.client.gocql.protoversion"
	Timeout                  = "transport.client.gocql.timeout"
	ConnectTimeout           = "transport.client.gocql.connecttimeout"
	Keyspace                 = "transport.client.gocql.keyspace"
	NumConns                 = "transport.client.gocql.numconns"
	Consistency              = "transport.client.gocql.consistency"
	SocketKeepalive          = "transport.client.gocql.socketkeepalive"
	MaxPreparedStmts         = "transport.client.gocql.maxpreparedstmts"
	MaxRoutingKeyInfo        = "transport.client.gocql.maxroutingkeyinfo"
	PageSize                 = "transport.client.gocql.pagesize"
	DefaultTimestamp         = "transport.client.gocql.defaulttimestamp"
	ReconnectInterval        = "transport.client.gocql.reconnectinterval"
	MaxWaitSchemaAgreement   = "transport.client.gocql.maxwaitschemaagreement"
	DisableInitialHostLookup = "transport.client.gocql.disableinitialhostlookup"
	WriteCoalesceWaitTime    = "transport.client.gocql.writecoalescewaittime"
	HealthEnabled            = "transport.client.gocql.health.enabled"
	HealthDescription        = "transport.client.gocql.health.description"
	HealthRequired           = "transport.client.gocql.health.required"
)

func init() {

	log.Println("getting configurations for mongodb")

	config.Add(Hosts, []string{"127.0.0.1"}, "addresses for the initial connections")
	config.Add(Port, 9042, "define port")
	config.Add(Username, "", "define username")
	config.Add(Password, "", "define password")
	config.Add(CQLVersion, "3.0.0", "define cql version")
	config.Add(ProtoVersion, 0, "define version of the native protocol to use")
	config.Add(Timeout, 600*time.Millisecond, "connection timeout")
	config.Add(ConnectTimeout, 600*time.Millisecond, "initial connection timeout, used during initial dial to server")
	config.Add(Keyspace, "", "initial keyspace (optional)")
	config.Add(NumConns, 2, "number of connections per host")
	config.Add(Consistency, "QUORUM", "default consistency level (default: Quorum) (values: ANY, ONE, TWO, THREE, QUORUM, ALL, LOCAL_QUORUM, EACH_QUORUM, LOCAL_ONE)")
	config.Add(SocketKeepalive, 0*time.Millisecond, "The keepalive period to use, enabled if > 0 (default: 0)")
	config.Add(MaxPreparedStmts, 1000, "Sets the maximum cache size for prepared statements globally for gocql")
	config.Add(MaxRoutingKeyInfo, 1000, "Sets the maximum cache size for query info about statements for each session")
	config.Add(PageSize, 5000, "Default page size to use for created sessions")
	config.Add(DefaultTimestamp, true, "Sends a client side timestamp for all requests which overrides the timestamp at which it arrives at the server. (default: true, only enabled for protocol 3 and above)")
	config.Add(ReconnectInterval, 10*time.Millisecond, "If not zero, gocql attempt to reconnect known DOWN nodes in every ReconnectInterval")
	config.Add(MaxWaitSchemaAgreement, 60*time.Second, "The maximum amount of time to wait for schema agreement in a cluster after receiving a schema change frame")
	config.Add(DisableInitialHostLookup, true, "If true then the driver will not attempt to get host info from the system.peers table")
	config.Add(WriteCoalesceWaitTime, 200*time.Microsecond, "The time to wait for frames before flushing the frames connection to Cassandra")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, true, "define health description")

}
