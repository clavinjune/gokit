package slogutil

import (
	"strings"

	"golang.org/x/exp/slog"
)

type Writer struct {
	_      struct{}
	logger *slog.Logger
	level  slog.Level
}

func NewWriter(logger *slog.Logger, level slog.Level) *Writer {
	return &Writer{
		logger: logger,
		level:  level,
	}
}

func (s *Writer) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(strings.Trim(string(p), "\n"))
	s.logger.LogAttrs(s.level, msg)
	return len(msg), nil
}
