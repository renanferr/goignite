package giredis

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConfigRoot         = "gi.redis"
	Password           = ConfigRoot + ".password"
	MaxRetries         = ConfigRoot + ".maxRetries"
	MinRetryBackoff    = ConfigRoot + ".minRetryBackoff"
	MaxRetryBackoff    = ConfigRoot + ".maxRetryBackoff"
	DialTimeout        = ConfigRoot + ".dialTimeout"
	ReadTimeout        = ConfigRoot + ".readTimeout"
	WriteTimeout       = ConfigRoot + ".writeTimeout"
	PoolSize           = ConfigRoot + ".poolSize"
	MinIdleConns       = ConfigRoot + ".minIdleConns"
	MaxConnAge         = ConfigRoot + ".maxConnAge"
	PoolTimeout        = ConfigRoot + ".poolTimeout"
	IdleTimeout        = ConfigRoot + ".idleTimeout"
	IdleCheckFrequency = ConfigRoot + ".idleCheckFrequency"
	Addr               = ConfigRoot + ".client.addr"
	Network            = ConfigRoot + ".client.network"
	DB                 = ConfigRoot + ".client.db"
	SentinelMaster     = ConfigRoot + ".sentinel.masterName"
	SentinelAddr       = ConfigRoot + ".sentinel.addrs"
	SentinelPassword   = ConfigRoot + ".sentinel.password"
	Addrs              = ConfigRoot + ".cluster.addrs"
	MaxRedirects       = ConfigRoot + ".cluster.maxRedirects"
	ReadOnly           = ConfigRoot + ".cluster.readOnly"
	RouteByLatency     = ConfigRoot + ".cluster.routeByLatency"
	RouteRandomly      = ConfigRoot + ".cluster.routeRandomly"
)

func init() {

	giconfig.Add(Addrs, []string{"127.0.0.1:6379"}, "a seed list of host:port addresses of cluster nodes")
	giconfig.Add(MaxRedirects, 8, "the maximum number of retries before giving up")
	giconfig.Add(ReadOnly, false, "enables read-only commands on slave nodes")
	giconfig.Add(RouteByLatency, false, "allows routing read-only commands to the closest master or slave node")
	giconfig.Add(RouteRandomly, false, "allows routing read-only commands to the random master or slave node")
	giconfig.Add(Password, "", "optional password. Must match the password specified in the requirepass server configuration option")
	giconfig.Add(MaxRetries, 0, "maximum number of retries before giving up")
	giconfig.Add(MinRetryBackoff, 8*time.Millisecond, "minimum backoff between each retry")
	giconfig.Add(MaxRetryBackoff, 512*time.Millisecond, "maximum backoff between each retry")
	giconfig.Add(DialTimeout, 5*time.Second, "dial timeout for establishing new connections")
	giconfig.Add(ReadTimeout, 3*time.Second, "timeout for socket reads. If reached, commands will fail with a timeout instead of blocking. Use value -1 for no timeout and 0 for default")
	giconfig.Add(WriteTimeout, 3*time.Second, "timeout for socket writes. If reached, commands will fail")
	giconfig.Add(PoolSize, 10, "maximum number of socket connections")
	giconfig.Add(MinIdleConns, 2, "minimum number of idle connections which is useful when establishing new connection is slow")
	giconfig.Add(MaxConnAge, 0*time.Millisecond, "connection age at which client retires (closes) the connection")
	giconfig.Add(PoolTimeout, 4*time.Second, "amount of time client waits for connection if all connections are busy before returning an error")
	giconfig.Add(IdleTimeout, 5*time.Minute, "amount of time after which client closes idle connections. Should be less than server's timeout")
	giconfig.Add(IdleCheckFrequency, 1*time.Minute, "frequency of idle checks made by idle connections reaper. Default is 1 minute. -1 disables idle connections reaper, but idle connections are still discarded by the client if IdleTimeout is set")
	giconfig.Add(Addr, "127.0.0.1:6379", "host:port address")
	giconfig.Add(Network, "tcp", "the network type, either tcp or unix")
	giconfig.Add(DB, 0, "database to be selected after connecting to the server")
	giconfig.Add(SentinelMaster, "", "redis sentinel master name")
	giconfig.Add(SentinelAddr, nil, "redis sentinel addr list host:port")
	giconfig.Add(SentinelPassword, "", "redis sentinel password")
}
