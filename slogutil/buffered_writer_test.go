package slogutil_test

import (
	"bufio"
	"bytes"
	"testing"
	"time"

	"github.com/clavinjune/gokit/slogutil"
	"github.com/stretchr/testify/require"
)

func TestBufferedWriter(t *testing.T) {
	r := require.New(t)
	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	w := slogutil.NewBufferedWriter(&slogutil.BufferedWriterOption{
		Writer:        bw,
		FlushInterval: time.Second,
		BufferSize:    10, // 10 Bytes
	})

	w.Write([]byte("foobar"))
	r.Equal("", b.String())
	time.Sleep(2 * time.Second)
	r.Equal("foobar", b.String())
	b.Reset()

	w.Write([]byte("10 bytes of data should be flushed right away"))
	r.Equal("10 bytes of data should be flushed right away", b.String())
	b.Reset()

	w.Write([]byte("xyz"))
	w.Flush()
	r.Equal("xyz", b.String())
}
