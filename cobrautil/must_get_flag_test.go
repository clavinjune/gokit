package cobrautil_test

import (
	"os"
	"testing"

	"github.com/clavinjune/gokit/cobrautil"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Setenv("TESTING_MUST_GET_BOOL_FLAG", "false")
	os.Setenv("TESTING_MUST_GET_INT_FLAG", "333")
	os.Setenv("TESTING_MUST_GET_STRING_FLAG", "fetched_from_env")
	code := m.Run()
	os.Clearenv()

	os.Exit(code)
}

func TestMustGetBoolFlag(t *testing.T) {
	tt := []struct {
		_                struct{}
		Name             string
		FlagDefaultValue bool
		CheckEnv         bool
		Expected         bool
	}{
		{
			Name:             "default",
			FlagDefaultValue: true,
			CheckEnv:         false,
			Expected:         true,
		},
		{
			Name:             "checkEnv",
			FlagDefaultValue: true,
			CheckEnv:         true,
			Expected:         false,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			cmd := &cobra.Command{Use: "testing"}
			cmd.Flags().Bool("must-get-bool-flag", tc.FlagDefaultValue, "test flag")

			r.Equal(tc.Expected, cobrautil.MustGetBoolFlag(cmd, "must-get-bool-flag", tc.CheckEnv))
			r.Panics(func() {
				cobrautil.MustGetBoolFlag(cmd, "unknown")
			})
		})
	}
}

func TestMustGetIntFlag(t *testing.T) {
	tt := []struct {
		_                struct{}
		Name             string
		FlagDefaultValue int
		CheckEnv         bool
		Expected         int
	}{
		{
			Name:             "default",
			FlagDefaultValue: 1,
			CheckEnv:         false,
			Expected:         1,
		},
		{
			Name:             "checkEnv",
			FlagDefaultValue: 1,
			CheckEnv:         true,
			Expected:         333,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			cmd := &cobra.Command{Use: "testing"}
			cmd.Flags().Int("must-get-int-flag", tc.FlagDefaultValue, "test flag")

			r.Equal(tc.Expected, cobrautil.MustGetIntFlag(cmd, "must-get-int-flag", tc.CheckEnv))
			r.Panics(func() {
				cobrautil.MustGetIntFlag(cmd, "unknown")
			})
		})
	}
}

func TestMustGetStringFlag(t *testing.T) {
	tt := []struct {
		_                struct{}
		Name             string
		FlagDefaultValue string
		CheckEnv         bool
		Expected         string
	}{
		{
			Name:             "default",
			FlagDefaultValue: "default",
			CheckEnv:         false,
			Expected:         "default",
		},
		{
			Name:             "checkEnv",
			FlagDefaultValue: "default",
			CheckEnv:         true,
			Expected:         "fetched_from_env",
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			cmd := &cobra.Command{Use: "testing"}
			cmd.Flags().String("must-get-string-flag", tc.FlagDefaultValue, "test flag")

			r.Equal(tc.Expected, cobrautil.MustGetStringFlag(cmd, "must-get-string-flag", tc.CheckEnv))
			r.Panics(func() {
				cobrautil.MustGetStringFlag(cmd, "unknown")
			})
		})
	}
}

func TestCheckEnv(t *testing.T) {
	tt := []struct {
		_             struct{}
		Name          string
		FlagValue     string
		ExpectedValue int
		ExpectedPanic bool
	}{
		{
			Name:          "with env value",
			FlagValue:     "123456",
			ExpectedValue: 123456,
			ExpectedPanic: false,
		},
		{
			Name:          "no env value",
			FlagValue:     "",
			ExpectedValue: 123,
			ExpectedPanic: false,
		},
		{
			Name:          "different type env value",
			FlagValue:     "qweqwe",
			ExpectedValue: 123,
			ExpectedPanic: true,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			r := require.New(t)

			cmd := &cobra.Command{Use: "testing"}
			cmd.Flags().Int("check-env", 123, "test flag")

			os.Setenv("TESTING_CHECK_ENV", tc.FlagValue)

			actual := 123
			fn := func() {
				actual = cobrautil.MustGetIntFlag(cmd, "check-env", true)
			}
			if tc.ExpectedPanic {
				r.Panics(fn)
			} else {
				r.NotPanics(fn)
			}
			r.Equal(tc.ExpectedValue, actual)
		})
	}
}
