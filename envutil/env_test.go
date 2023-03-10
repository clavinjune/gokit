package envutil_test

import (
	"os"
	"testing"

	"github.com/clavinjune/gokit/stringutil"

	"github.com/clavinjune/gokit/envutil"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	defer os.Clearenv()
	m.Run()
}

func TestGet(t *testing.T) {
	tt := []struct {
		_             struct{}
		Name          string
		ExpectedValue string
		ExpectedError error
		Opt           *envutil.Option
		EnvValue      string
	}{
		{
			Name:          "default option - empty value",
			ExpectedValue: "",
			ExpectedError: nil,
			Opt:           nil,
			EnvValue:      "",
		},
		{
			Name:          "default option - non empty value",
			ExpectedValue: "real value",
			ExpectedError: nil,
			EnvValue:      "real value",
		},
		{
			Name:          "is required - empty value",
			ExpectedValue: "",
			ExpectedError: envutil.ErrRequired,
			Opt: &envutil.Option{
				IsRequired: true,
			},
			EnvValue: "",
		},
		{
			Name:          "default value - empty value",
			ExpectedValue: "default value",
			ExpectedError: nil,
			Opt: &envutil.Option{
				DefaultValue: "default value",
			},
			EnvValue: "",
		},
		{
			Name:          "default value - non empty value",
			ExpectedValue: "real value",
			ExpectedError: nil,
			Opt: &envutil.Option{
				DefaultValue: "default value",
			},
			EnvValue: "real value",
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)
			key := "gokit" + stringutil.Random(5)

			if tc.EnvValue != "" {
				os.Setenv(key, tc.EnvValue)
				defer os.Unsetenv(key)
			}

			actualValue, actualErr := envutil.Get(key, tc.Opt)
			r.Equal(tc.ExpectedValue, actualValue)

			if tc.ExpectedError != nil {
				r.Error(actualErr)
				r.ErrorIs(actualErr, tc.ExpectedError)
			} else {
				r.NoError(actualErr)
			}
		})
	}
}
