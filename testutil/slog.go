package testutil

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/clavinjune/gokit/slogutil"
)

func NewSlog(t *testing.T) (*slog.Logger, *bytes.Buffer) {
	t.Helper()
	var b bytes.Buffer
	return slogutil.New(&slogutil.Option{
		IsDebug: true,
		Writer:  &b,
	}).WithGroup("test"), &b
}
