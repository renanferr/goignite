package http

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
)

type Options struct {
	MaxIdleConnPerHost    int           `config:"maxidleconnperhost"`
	MaxIdleConn           int           `config:"maxidleconn"`
	MaxConnsPerHost       int           `config:"maxconnsperhost"`
	IdleConnTimeout       time.Duration `config:"idleconntimeout"`
	DisableKeepAlives     bool          `config:"disablekeepalives"`
	HTTP2                 bool          `config:"http2"`
	TLSHandshakeTimeout   time.Duration `config:"tlshandshaketimeout"`
	Timeout               time.Duration
	KeepAlive             time.Duration `config:"keepalive"`
	ExpectContinueTimeout time.Duration `config:"expectcontinuetimeout"`
	DualStack             bool          `config:"dualstack"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(Pkg, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
