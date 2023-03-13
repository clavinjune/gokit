package cobrautil

import "github.com/clavinjune/gokit/slogutil"

var (
	DefaultOption = Option{
		SlogOption:   slogutil.DefaultOption,
		SetOutToSlog: false,
		SetErrToSlog: false,
	}
)

type Option struct {
	_            struct{}
	SlogOption   slogutil.Option
	SetOutToSlog bool
	SetErrToSlog bool
}
