package slogutil

import (
	"io"
	"log/slog"
	"os"
	"time"
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

	DefaultBufferedWriterOption = BufferedWriterOption{
		Writer:        os.Stdout,
		BufferSize:    30 * 1024, // 30KB
		FlushInterval: 15 * time.Second,
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

func (o *Option) WriterOrStdout() io.Writer {
	if o.Writer != nil {
		return o.Writer
	}

	return os.Stdout
}

type BufferedWriterOption struct {
	_             struct{}
	Writer        io.Writer
	BufferSize    int
	FlushInterval time.Duration
}

func (b *BufferedWriterOption) MustWriter() io.Writer {
	if b.Writer != nil {
		return b.Writer
	}

	return DefaultBufferedWriterOption.Writer
}

func (b *BufferedWriterOption) MustBufferSize() int {
	if b.BufferSize != 0 {
		return b.BufferSize
	}

	return DefaultBufferedWriterOption.BufferSize
}

func (b *BufferedWriterOption) MustFlushInterval() time.Duration {
	if b.FlushInterval > 0 {
		return b.FlushInterval
	}

	return DefaultBufferedWriterOption.FlushInterval
}
