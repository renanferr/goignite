package buntdb

import (
	"github.com/b2wdigital/goignite/pkg/config"
)

type Options struct {
	Path       string
	SyncPolicy int `config:"syncpolicy"`
	AutoShrink struct {
		Percentage int
		MinSize    int
		Disabled   bool
	} `config:"autoshrink"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.buntdb", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
