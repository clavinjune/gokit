package slogutil

import (
	"context"
	"strings"

	"golang.org/x/exp/slog"
)

type Writer struct {
	_      struct{}
	ctx    context.Context
	logger *slog.Logger
	level  slog.Level
}

func NewWriter(ctx context.Context, logger *slog.Logger, level slog.Level) *Writer {
	return &Writer{
		ctx:    ctx,
		logger: logger,
		level:  level,
	}
}

func (s *Writer) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(strings.Trim(string(p), "\n"))
	s.logger.LogAttrs(s.ctx, s.level, msg)
	return len(msg), nil
}
