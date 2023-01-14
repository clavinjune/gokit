package cobrautil_test

import (
	"testing"

	"github.com/clavinjune/gokit/cobrautil"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestMustGetBoolFlag(t *testing.T) {
	r := require.New(t)
	cmd := &cobra.Command{}
	cmd.Flags().Bool("foo", true, "test flag")

	r.True(cobrautil.MustGetBoolFlag(cmd, "foo"))
	r.Panics(func() {
		cobrautil.MustGetBoolFlag(cmd, "unknown")
	})
}
