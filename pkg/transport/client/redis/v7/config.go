package redis

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Password           = "transport.client.redis.password"
	MaxRetries         = "transport.client.redis.maxretries"
	MinRetryBackoff    = "transport.client.redis.minretrybackoff"
	MaxRetryBackoff    = "transport.client.redis.maxretrybackoff"
	DialTimeout        = "transport.client.redis.dialtimeout"
	ReadTimeout        = "transport.client.redis.readtimeout"
	WriteTimeout       = "transport.client.redis.writetimeout"
	PoolSize           = "transport.client.redis.poolsize"
	MinIdleConns       = "transport.client.redis.minidleconns"
	MaxConnAge         = "transport.client.redis.maxconnage"
	PoolTimeout        = "transport.client.redis.pooltimeout"
	IdleTimeout        = "transport.client.redis.idletimeout"
	IdleCheckFrequency = "transport.client.redis.idlecheckfrequency"
	Addr               = "transport.client.redis.client.addr"
	Network            = "transport.client.redis.client.network"
	DB                 = "transport.client.redis.client.db"
	Addrs              = "transport.client.redis.cluster.addrs"
	MaxRedirects       = "transport.client.redis.cluster.maxredirects"
	ReadOnly           = "transport.client.redis.cluster.readonly"
	RouteByLatency     = "transport.client.redis.cluster.routebylatency"
	RouteRandomly      = "transport.client.redis.cluster.routerandomly"
	HealthEnabled      = "transport.client.redis.health.enabled"
	HealthDescription  = "transport.client.redis.health.description"
	HealthRequired     = "transport.client.redis.health.required"
	NewRelicEnabled    = "transport.client.redis.newrelic.enabled"
)

func init() {

	log.Println("getting configurations for redis")

	config.Add(Addrs, []string{"127.0.0.1:6379"}, "a seed list of host:port addresses of cluster nodes")
	config.Add(MaxRedirects, 8, "the maximum number of retries before giving up")
	config.Add(ReadOnly, false, "enables read-only commands on slave nodes")
	config.Add(RouteByLatency, false, "allows routing read-only commands to the closest master or slave node")
	config.Add(RouteRandomly, false, "allows routing read-only commands to the random master or slave node")
	config.Add(Password, "", "optional password. Must match the password specified in the requirepass server configuration option")
	config.Add(MaxRetries, 0, "maximum number of retries before giving up")
	config.Add(MinRetryBackoff, 8*time.Millisecond, "minimum backoff between each retry")
	config.Add(MaxRetryBackoff, 512*time.Millisecond, "maximum backoff between each retry")
	config.Add(DialTimeout, 5*time.Second, "dial timeout for establishing new connections")
	config.Add(ReadTimeout, 3*time.Second, "timeout for socket reads. If reached, commands will fail with a timeout instead of blocking. Use value -1 for no timeout and 0 for default")
	config.Add(WriteTimeout, 3*time.Second, "timeout for socket writes. If reached, commands will fail")
	config.Add(PoolSize, 10, "maximum number of socket connections")
	config.Add(MinIdleConns, 2, "minimum number of idle connections which is useful when establishing new connection is slow")
	config.Add(MaxConnAge, 0*time.Millisecond, "connection age at which client retires (closes) the connection")
	config.Add(PoolTimeout, 4*time.Second, "amount of time client waits for connection if all connections are busy before returning an error")
	config.Add(IdleTimeout, 5*time.Minute, "amount of time after which client closes idle connections. Should be less than server's timeout")
	config.Add(IdleCheckFrequency, 1*time.Minute, "frequency of idle checks made by idle connections reaper. Default is 1 minute. -1 disables idle connections reaper, but idle connections are still discarded by the client if IdleTimeout is set")
	config.Add(Addr, "127.0.0.1:6379", "host:port address")
	config.Add(Network, "tcp", "the network type, either tcp or unix")
	config.Add(DB, 0, "database to be selected after connecting to the server")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, true, "define health description")
	config.Add(NewRelicEnabled, false, "enable/disable newrelic")
}
