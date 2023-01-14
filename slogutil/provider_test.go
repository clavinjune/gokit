package slogutil_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/clavinjune/gokit/slogutil"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
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

			log := slogutil.New(&tc.Option)
			log.Info("info")
			log.Debug("debug")

			r.Equal(tc.Expected, buf.String())
		})
	}
}

func TestNew_ShortFile(t *testing.T) {
	buf := new(bytes.Buffer)
	opt := &slogutil.Option{
		IsDebug:      true,
		IsJSON:       true,
		IsShortFile:  false,
		RedactedKeys: []string{"time"},
		Writer:       buf,
	}
	slogutil.New(opt).Info("long")

	opt.IsShortFile = true
	slogutil.New(opt).Info("short")

	r := require.New(t)
	m := map[string]any{}

	dec := json.NewDecoder(buf)
	r.NoError(dec.Decode(&m))
	r.Equal(m["msg"], "long")
	r.Contains(m["source"], "slogutil/provider_test.go:")

	r.NoError(dec.Decode(&m))
	r.Equal(m["msg"], "short")
	r.NotContains(m["source"], "/")
	r.Contains(m["source"], "provider_test.go:")
}
