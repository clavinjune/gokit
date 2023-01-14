package stringutil_test

import (
	"testing"

	"github.com/clavinjune/gokit/internal/stringutil"
	"github.com/stretchr/testify/require"
)

func TestRandom(t *testing.T) {
	tt := []struct {
		_                   struct{}
		Name                string
		N                   int
		ExpectedValueLength int
	}{
		{
			Name:                "less than zero",
			N:                   -1,
			ExpectedValueLength: 0,
		},
		{
			Name:                "zero length",
			N:                   0,
			ExpectedValueLength: 0,
		},
		{
			Name:                "positive length",
			N:                   10,
			ExpectedValueLength: 10,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := len(stringutil.Random(tc.N))
			r.Equal(tc.ExpectedValueLength, actual)
		})
	}
}
