package argutil_test

import (
	"testing"

	"github.com/clavinjune/gokit/internal/argutil"
	"github.com/stretchr/testify/require"
)

func TestFirstOrDefault(t *testing.T) {
	tt := []struct {
		_          struct{}
		Name       string
		DefaultArg string
		Args       []string
		Expected   string
	}{
		{
			Name:       "nil",
			DefaultArg: "default",
			Args:       nil,
			Expected:   "default",
		},
		{
			Name:       "zero length",
			DefaultArg: "default",
			Args:       []string{},
			Expected:   "default",
		},
		{
			Name:       "1 length",
			DefaultArg: "default",
			Args:       []string{"first"},
			Expected:   "first",
		},
		{
			Name:       "2 length",
			DefaultArg: "default",
			Args:       []string{"first", "second"},
			Expected:   "first",
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := argutil.FirstOrDefault(tc.DefaultArg, tc.Args...)
			r.Equal(tc.Expected, actual)
		})
	}
}
