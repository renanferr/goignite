package freecache

import "github.com/b2wdigital/goignite/v2/config"

type Options struct {
	CacheSize int
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
