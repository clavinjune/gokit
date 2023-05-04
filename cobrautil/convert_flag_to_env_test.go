package cobrautil

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestToEnvName(t *testing.T) {
	tt := []struct {
		_        struct{}
		Name     string
		CmdUse   string
		FlagName string
		Expected string
	}{
		{
			Name:     "no name",
			CmdUse:   "",
			FlagName: "foo-bar",
			Expected: "FOO_BAR",
		},
		{
			Name:     "no name double underscore",
			CmdUse:   "",
			FlagName: "foo--bar",
			Expected: "FOO_BAR",
		},
		{
			Name:     "with name",
			CmdUse:   "gokit",
			FlagName: "foo-bar",
			Expected: "GOKIT_FOO_BAR",
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := toEnvName(&cobra.Command{Use: tc.CmdUse}, tc.FlagName)
			r.Equal(tc.Expected, actual)
		})
	}

}
