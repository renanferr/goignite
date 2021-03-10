package giecho

import giconfig "github.com/b2wdigital/goignite/v2/config"

type Options struct {
	HideBanner bool
	Port       int
	Json       struct {
		Pretty struct {
			Enabled bool
		}
	}
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
