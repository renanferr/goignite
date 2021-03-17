package buntdb

import "github.com/b2wdigital/goignite/v2/config"

type Options struct {
	Path       string
	SyncPolicy int
	AutoShrink struct {
		Percentage int
		MinSize    int
		Disabled   bool
	} `config:"autoShrink"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
