package envutil

import (
	"errors"
	"fmt"
)

var (
	// errBase identifies the errors come from gokit/envutil package
	errBase = errors.New("gokit/envutil")

	// ErrRequired identifies the envutil is required
	ErrRequired = fmt.Errorf("%w: envutil is required", errBase)
)

// errRequiredFmt formats ErrRequired error with the required key
func errRequiredFmt(key string) error {
	return fmt.Errorf("%w: %+q", ErrRequired, key)
}
