package config

import (
	"os"
)

// Get gets value from os.Getenv with customized behavior from *Option
func Get(key string, opts ...*Option) (string, error) {
	opt := option(opts...)
	value := os.Getenv(key)

	if value == "" && opt.DefaultValue != "" {
		value = opt.DefaultValue
	}

	if value == "" && opt.IsRequired {
		return "", errRequiredFmt(key)
	}

	return value, nil
}

// option use the first Option if not empty
// otherwise use DefaultOption
func option(opts ...*Option) *Option {
	var opt *Option
	if len(opts) >= 1 && opts[0] != nil {
		opt = opts[0]
	} else {
		opt = &DefaultOption
	}

	return opt
}
