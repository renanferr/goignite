package client

import (
	"time"

	"github.com/b2wdigital/goignite/v2/config"
)

type Options struct {
	MaxIdleConnPerHost    int
	MaxIdleConn           int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	DisableKeepAlives     bool
	DisableCompression    bool
	ForceHTTP2            bool          `config:"forceHTTP2"`
	TLSHandshakeTimeout   time.Duration `config:"TLSHandshakeTimeout"`
	Timeout               time.Duration
	KeepAlive             time.Duration
	ExpectContinueTimeout time.Duration
	DualStack             bool
	DialTimeout           time.Duration
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
