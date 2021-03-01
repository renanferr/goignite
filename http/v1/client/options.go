package gihttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

type Options struct {
	MaxIdleConnPerHost    int
	MaxIdleConn           int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	DisableKeepAlives     bool
	ForceHTTP2            bool          `config:"forceHTTP2"`
	TLSHandshakeTimeout   time.Duration `config:"TLSHandshakeTimeout"`
	Timeout               time.Duration
	KeepAlive             time.Duration
	ExpectContinueTimeout time.Duration
	DualStack             bool
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
