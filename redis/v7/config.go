package giredis

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root               = "gi.redis"
	password           = root + ".password"
	maxRetries         = root + ".maxRetries"
	minRetryBackoff    = root + ".minRetryBackoff"
	maxRetryBackoff    = root + ".maxRetryBackoff"
	dialTimeout        = root + ".dialTimeout"
	readTimeout        = root + ".readTimeout"
	writeTimeout       = root + ".writeTimeout"
	poolSize           = root + ".poolSize"
	minIdleConns       = root + ".minIdleConns"
	maxConnAge         = root + ".maxConnAge"
	poolTimeout        = root + ".poolTimeout"
	idleTimeout        = root + ".idleTimeout"
	idleCheckFrequency = root + ".idleCheckFrequency"
	addr               = root + ".client.addr"
	network            = root + ".client.network"
	db                 = root + ".client.db"
	sentinelMaster     = root + ".sentinel.masterName"
	sentinelAddr       = root + ".sentinel.addrs"
	sentinelPassword   = root + ".sentinel.password"
	addrs              = root + ".cluster.addrs"
	maxRedirects       = root + ".cluster.maxRedirects"
	readOnly           = root + ".cluster.readOnly"
	routeByLatency     = root + ".cluster.routeByLatency"
	routeRandomly      = root + ".cluster.routeRandomly"
	ExtRoot            = root + ".Ext"
)

func init() {

	giconfig.Add(addrs, []string{"127.0.0.1:6379"}, "a seed list of host:port addresses of cluster nodes")
	giconfig.Add(maxRedirects, 8, "the maximum number of retries before giving up")
	giconfig.Add(readOnly, false, "enables read-only commands on slave nodes")
	giconfig.Add(routeByLatency, false, "allows routing read-only commands to the closest master or slave node")
	giconfig.Add(routeRandomly, false, "allows routing read-only commands to the random master or slave node")
	giconfig.Add(password, "", "optional password. Must match the password specified in the requirepass server configuration option")
	giconfig.Add(maxRetries, 0, "maximum number of retries before giving up")
	giconfig.Add(minRetryBackoff, 8*time.Millisecond, "minimum backoff between each retry")
	giconfig.Add(maxRetryBackoff, 512*time.Millisecond, "maximum backoff between each retry")
	giconfig.Add(dialTimeout, 5*time.Second, "dial timeout for establishing new connections")
	giconfig.Add(readTimeout, 3*time.Second, "timeout for socket reads. If reached, commands will fail with a timeout instead of blocking. Use value -1 for no timeout and 0 for default")
	giconfig.Add(writeTimeout, 3*time.Second, "timeout for socket writes. If reached, commands will fail")
	giconfig.Add(poolSize, 10, "maximum number of socket connections")
	giconfig.Add(minIdleConns, 2, "minimum number of idle connections which is useful when establishing new connection is slow")
	giconfig.Add(maxConnAge, 0*time.Millisecond, "connection age at which client retires (closes) the connection")
	giconfig.Add(poolTimeout, 4*time.Second, "amount of time client waits for connection if all connections are busy before returning an error")
	giconfig.Add(idleTimeout, 5*time.Minute, "amount of time after which client closes idle connections. Should be less than server's timeout")
	giconfig.Add(idleCheckFrequency, 1*time.Minute, "frequency of idle checks made by idle connections reaper. Default is 1 minute. -1 disables idle connections reaper, but idle connections are still discarded by the client if idleTimeout is set")
	giconfig.Add(addr, "127.0.0.1:6379", "host:port address")
	giconfig.Add(network, "tcp", "the network type, either tcp or unix")
	giconfig.Add(db, 0, "database to be selected after connecting to the server")
	giconfig.Add(sentinelMaster, "", "redis sentinel master name")
	giconfig.Add(sentinelAddr, nil, "redis sentinel addr list host:port")
	giconfig.Add(sentinelPassword, "", "redis sentinel password")
}
