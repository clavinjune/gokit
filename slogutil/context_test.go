package slogutil_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/clavinjune/gokit/slogutil"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	slogutil.New(&slogutil.DefaultOption)
	m.Run()
}

func TestFrom(t *testing.T) {
	customized := slog.Default().With(
		slog.Bool("default", false),
	)

	tt := []struct {
		Name     string
		Context  context.Context
		Expected *slog.Logger
	}{
		{
			Name:     "from empty context, should return default",
			Context:  context.Background(),
			Expected: slog.Default(),
		},
		{
			Name:     "from filled context, should not return default",
			Context:  slogutil.Put(context.Background(), customized),
			Expected: customized,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := slogutil.From(tc.Context)
			r.Equal(tc.Expected, actual)
		})
	}
}
