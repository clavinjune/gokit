package slogutil

import (
	"context"
	"log/slog"
)

type key struct{}

var keyInst *key = &key{}

func From(ctx context.Context) *slog.Logger {
	val, ok := ctx.Value(keyInst).(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	return val
}

func Put(ctx context.Context, log *slog.Logger) context.Context {
	return context.WithValue(ctx, keyInst, log)
}
