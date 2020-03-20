package gocql

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Hosts                    []string
	Port                     int
	Username                 string
	Password                 string
	CQLVersion               string `config:"cqlversion"`
	ProtoVersion             int    `config:"protoversion"`
	Timeout                  time.Duration
	ConnectTimeout           time.Duration `config:"connecttimeout"`
	Keyspace                 string
	NumConns                 int `config:"numconns"`
	Consistency              string
	SocketKeepalive          time.Duration `config:"socketkeepalive"`
	MaxPreparedStmts         int           `config:"maxpreparedstmts"`
	MaxRoutingKeyInfo        int           `config:"maxroutingkeyinfo"`
	PageSize                 int           `config:"pagesize"`
	DefaultTimestamp         bool          `config:"defaulttimestamp"`
	ReconnectInterval        time.Duration `config:"reconnectinterval"`
	MaxWaitSchemaAgreement   time.Duration `config:"maxwaitschemaagreement"`
	DisableInitialHostLookup bool          `config:"disableinitialhostlookup"`
	WriteCoalesceWaitTime    time.Duration `config:"writecoalescewaittime"`
	Health                   OptionsHealth
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Hosts(value []string) optionsBuilder {
	return builder.Set(b, "Hosts", value).(optionsBuilder)
}

func (b optionsBuilder) Port(value int) optionsBuilder {
	return builder.Set(b, "Port", value).(optionsBuilder)
}

func (b optionsBuilder) Username(value int) optionsBuilder {
	return builder.Set(b, "Username", value).(optionsBuilder)
}

func (b optionsBuilder) Password(value int) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) CQLVersion(value string) optionsBuilder {
	return builder.Set(b, "CQLVersion", value).(optionsBuilder)
}

func (b optionsBuilder) ProtoVersion(value int) optionsBuilder {
	return builder.Set(b, "ProtoVersion", value).(optionsBuilder)
}

func (b optionsBuilder) Timeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "Timeout", value).(optionsBuilder)
}

func (b optionsBuilder) ConnectTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "ConnectTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) Keyspace(value string) optionsBuilder {
	return builder.Set(b, "Keyspace", value).(optionsBuilder)
}

func (b optionsBuilder) NumConns(value int) optionsBuilder {
	return builder.Set(b, "NumConns", value).(optionsBuilder)
}

func (b optionsBuilder) Consistency(value string) optionsBuilder {
	return builder.Set(b, "Consistency", value).(optionsBuilder)
}

func (b optionsBuilder) SocketKeepalive(value time.Duration) optionsBuilder {
	return builder.Set(b, "SocketKeepalive", value).(optionsBuilder)
}

func (b optionsBuilder) MaxPreparedStmts(value int) optionsBuilder {
	return builder.Set(b, "MaxPreparedStmts", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRoutingKeyInfo(value int) optionsBuilder {
	return builder.Set(b, "MaxRoutingKeyInfo", value).(optionsBuilder)
}

func (b optionsBuilder) PageSize(value int) optionsBuilder {
	return builder.Set(b, "PageSize", value).(optionsBuilder)
}

func (b optionsBuilder) DefaultTimestamp(value bool) optionsBuilder {
	return builder.Set(b, "DefaultTimestamp", value).(optionsBuilder)
}

func (b optionsBuilder) ReconnectInterval(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReconnectInterval", value).(optionsBuilder)
}

func (b optionsBuilder) MaxWaitSchemaAgreement(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxWaitSchemaAgreement", value).(optionsBuilder)
}

func (b optionsBuilder) DisableInitialHostLookup(value bool) optionsBuilder {
	return builder.Set(b, "DisableInitialHostLookup", value).(optionsBuilder)
}

func (b optionsBuilder) WriteCoalesceWaitTime(value time.Duration) optionsBuilder {
	return builder.Set(b, "WriteCoalesceWaitTime", value).(optionsBuilder)
}

func (b optionsBuilder) Health(health OptionsHealth) optionsBuilder {
	return builder.Set(b, "Health", health).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.gocql", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
