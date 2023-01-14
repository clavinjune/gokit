package config

import (
	"os"

	"github.com/clavinjune/gokit/argutil"
)

// Get gets value from os.Getenv with customized behavior from *Option
func Get(key string, opts ...*Option) (string, error) {
	opt := argutil.FirstOrDefault(&DefaultOption, opts...)
	value := os.Getenv(key)

	if value == "" && opt.DefaultValue != "" {
		value = opt.DefaultValue
	}

	if value == "" && opt.IsRequired {
		return "", errRequiredFmt(key)
	}

	return value, nil
}
