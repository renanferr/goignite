package server

import "github.com/b2wdigital/goignite/v2/core/config"

type Options struct {
	Port                 int
	MaxConcurrentStreams int64
	TLS                  struct {
		Enabled  bool
		CertFile string
		KeyFile  string
		CAFile   string `config:"caFile"`
	} `config:"tls"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
