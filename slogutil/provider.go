package slogutil

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/clavinjune/gokit/argutil"
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
			switch a.Key {
			case slog.SourceKey:
				src := a.Value.Any().(*slog.Source)
				if opt.IsShortFile {
					a.Value = slog.StringValue(
						fmt.Sprintf("%s:%d", filepath.Base(src.File), src.Line),
					)
				} else {
					a.Value = slog.StringValue(
						fmt.Sprintf("%s:%d", src.File, src.Line),
					)
				}
			}

			if _, ok := redactedKeySet[a.Key]; ok {
				a.Value = RedactedValueAttr
			}

			return a
		},
	}

	var h slog.Handler
	if opt.IsJSON {
		h = slog.NewJSONHandler(opt.WriterOrStdout(), handlerOpt)
	} else {
		h = slog.NewTextHandler(opt.WriterOrStdout(), handlerOpt)
	}

	logger := slog.New(h)
	if opt.SetAsDefault {
		slog.SetDefault(logger)
	}

	return logger
}
