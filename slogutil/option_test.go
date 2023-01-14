package slogutil_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/clavinjune/gokit/slogutil"
	"github.com/stretchr/testify/require"
)

func TestOption_WriterOrStdout(t *testing.T) {
	r := require.New(t)
	opt := &slogutil.Option{}
	r.Equal(os.Stdout, opt.WriterOrStdout())

	buf := new(bytes.Buffer)
	opt.Writer = buf
	r.Equal(buf, opt.WriterOrStdout())
}
