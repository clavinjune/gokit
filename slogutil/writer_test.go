package slogutil_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/clavinjune/gokit/slogutil"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
)

func TestNewWriter(t *testing.T) {
	tt := []struct {
		_        struct{}
		Name     string
		Option   slogutil.Option
		Expected string
	}{
		{
			Name:     "key redaction",
			Option:   slogutil.DefaultOption,
			Expected: "time=[REDACTED] level=INFO msg=info\n",
		},
		{
			Name: "debug",
			Option: slogutil.Option{
				IsDebug:      true,
				RedactedKeys: []string{"source"},
			},
			Expected: `time=[REDACTED] level=INFO source=[REDACTED] msg=info
time=[REDACTED] level=DEBUG source=[REDACTED] msg=debug
`,
		},
		{
			Name: "json",
			Option: slogutil.Option{
				IsJSON: true,
			},
			Expected: `{"time":"[REDACTED]","level":"INFO","msg":"info"}
`,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			buf := new(bytes.Buffer)
			tc.Option.RedactedKeys = append(tc.Option.RedactedKeys, "time")
			tc.Option.Writer = buf

			logger := slogutil.New(&tc.Option)
			slogutil.NewWriter(context.Background(), logger, slog.LevelInfo).Write([]byte("info"))
			slogutil.NewWriter(context.Background(), logger, slog.LevelDebug).Write([]byte("debug"))

			r.Equal(tc.Expected, buf.String())
		})
	}
}
