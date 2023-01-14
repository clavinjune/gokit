package slogutil

import (
	"path/filepath"

	"github.com/clavinjune/gokit/argutil"

	"golang.org/x/exp/slog"
)

func New(opts ...*Option) *slog.Logger {
	opt := argutil.FirstOrDefault(&DefaultOption, opts...)

	level := slog.LevelInfo
	if opt.IsDebug {
		level = slog.LevelDebug
	}

	redactedKeySet := make(map[string]struct{})
	for _, key := range opt.RedactedKeys {
		redactedKeySet[key] = struct{}{}
	}

	handlerOpt := &slog.HandlerOptions{
		AddSource: opt.IsDebug,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey && opt.IsShortFile {
				a.Value = slog.StringValue(filepath.Base(a.Value.String()))
			}

			if _, ok := redactedKeySet[a.Key]; ok {
				a.Value = RedactedValueAttr
			}

			return a
		},
	}

	var h slog.Handler
	if opt.IsJSON {
		h = handlerOpt.NewJSONHandler(opt.WriterOrStdout())
	} else {
		h = handlerOpt.NewTextHandler(opt.WriterOrStdout())
	}

	logger := slog.New(h)
	if opt.SetAsDefault {
		slog.SetDefault(logger)
	}

	return logger
}
