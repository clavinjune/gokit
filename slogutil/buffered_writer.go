package slogutil

import (
	"bufio"
	"sync"
	"time"

	"github.com/clavinjune/gokit/argutil"
)

type BufferedWriter struct {
	sync.Mutex
	writer            *bufio.Writer
	bufferSize        int
	flushInterval     time.Duration
	currentBufferSize int
	lastFlush         time.Time
}

func NewBufferedWriter(opts ...*BufferedWriterOption) *BufferedWriter {
	opt := argutil.FirstOrDefault(&DefaultBufferedWriterOption, opts...)
	b := &BufferedWriter{
		writer:        bufio.NewWriter(opt.MustWriter()),
		bufferSize:    opt.MustBufferSize(),
		flushInterval: opt.MustFlushInterval(),
	}

	go b.autoFlush()
	return b
}

func (b *BufferedWriter) Write(p []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	n, err := b.writer.Write(p)
	if err != nil {
		return n, err
	}
	b.currentBufferSize += n

	if b.currentBufferSize >= b.bufferSize {
		err := b.Flush()
		if err != nil {
			return 0, err
		}
	}

	return n, err
}

func (b *BufferedWriter) Flush() error {
	if err := b.writer.Flush(); err != nil {
		return err
	}
	b.currentBufferSize = 0
	b.lastFlush = time.Now()

	return nil
}

func (b *BufferedWriter) autoFlush() {
	t := time.NewTicker(b.flushInterval)
	for range t.C {
		b.Lock()

		if time.Since(b.lastFlush) >= b.flushInterval {
			_ = b.Flush()
		}
		b.Unlock()
	}
}
