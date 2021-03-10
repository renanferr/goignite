package gihttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

type Options struct {
	Addr              string
	MaxHeaderBytes    int
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
