package slogutil

import (
	"io"
	"os"

	"golang.org/x/exp/slog"
)

const (
	RedactedValue = "[REDACTED]"
)

var (
	RedactedValueAttr = slog.StringValue(RedactedValue)

	DefaultOption = Option{
		SetAsDefault: true,
		IsDebug:      false,
		IsJSON:       false,
		IsShortFile:  true,
		RedactedKeys: []string{},
		Writer:       os.Stdout,
	}
)

type Option struct {
	_            struct{}
	SetAsDefault bool
	IsDebug      bool
	IsJSON       bool
	IsShortFile  bool
	RedactedKeys []string
	Writer       io.Writer
}
