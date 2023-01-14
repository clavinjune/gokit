package config

import (
	"errors"
	"fmt"
)

var (
	// errBase identifies the errors come from gokit/config package
	errBase error = errors.New("gokit/config")

	// ErrRequired identifies the config is required
	ErrRequired error = fmt.Errorf("%w: config is required", errBase)
)

// errRequiredFmt formats ErrRequired error with the required key
func errRequiredFmt(key string) error {
	return fmt.Errorf("%w: %+q", ErrRequired, key)
}
