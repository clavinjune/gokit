package cobrautil

import "github.com/clavinjune/gokit/slogutil"

var (
	DefaultOption = Option{
		SlogOption: slogutil.DefaultOption,
	}
)

type Option struct {
	_          struct{}
	SlogOption slogutil.Option
}
