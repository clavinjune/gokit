package cobrautil_test

import (
	"os"
	"testing"

	"github.com/clavinjune/gokit/cobrautil/v2"
	"github.com/clavinjune/gokit/testutil"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestMustBool(t *testing.T) {
	r := require.New(t)

	root := &cobra.Command{
		Use: "example",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cobrautil.MustBool(cmd, "debug"))
		},
	}

	t.Run("check usage", func(t *testing.T) {
		cobrautil.SetBool(root, "debug", "d", false, "debug")
		out, err := testutil.CobraExecute(t, root, "-h")
		r.NoError(err)
		r.Contains(out, "($EXAMPLE_DEBUG)")
		r.Contains(out, "-d")
		r.Contains(out, "--debug")
		root.ResetFlags()
	})

	t.Run("check value using flag", func(t *testing.T) {
		cobrautil.SetBool(root, "debug", "d", false, "debug")
		out, err := testutil.CobraExecute(t, root)
		r.NoError(err)
		r.Contains(out, "false")

		out, err = testutil.CobraExecute(t, root, "-d")
		r.NoError(err)
		r.Contains(out, "true")

		root.ResetFlags()
	})

	t.Run("check value using env var", func(t *testing.T) {
		cobrautil.SetBool(root, "debug", "d", false, "debug")
		out, err := testutil.CobraExecute(t, root)
		r.NoError(err)
		r.Contains(out, "false")

		os.Setenv("EXAMPLE_DEBUG", "true")
		defer os.Clearenv()
		out, err = testutil.CobraExecute(t, root)
		r.NoError(err)
		r.Contains(out, "true")

		root.ResetFlags()
	})
}

func TestMustString(t *testing.T) {
	r := require.New(t)

	root := &cobra.Command{
		Use: "example",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cobrautil.MustString(cmd, "name"))
		},
	}

	foo := &cobra.Command{
		Use: "foo",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cobrautil.MustString(cmd, "name"))
		},
	}

	bar := &cobra.Command{
		Use: "bar",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cobrautil.MustString(cmd, "name"))
		},
	}
	foo.AddCommand(bar)
	root.AddCommand(foo)

	t.Run("check usage", func(t *testing.T) {
		cobrautil.SetString(root, "name", "n", "john doe", "name")
		cobrautil.SetString(foo, "name", "n", "john doe", "name")
		cobrautil.SetString(bar, "name", "n", "john doe", "name")

		defer root.ResetFlags()
		defer foo.ResetFlags()
		defer bar.ResetFlags()

		out, err := testutil.CobraExecute(t, root, "-h")
		r.NoError(err)
		r.Contains(out, "($EXAMPLE_NAME)")
		r.Contains(out, "-n")
		r.Contains(out, "--name")

		out, err = testutil.CobraExecute(t, root, "foo", "-h")
		r.NoError(err)
		r.Contains(out, "($EXAMPLE_FOO_NAME)")
		r.Contains(out, "-n")
		r.Contains(out, "--name")

		out, err = testutil.CobraExecute(t, root, "foo", "bar", "-h")
		r.NoError(err)
		r.Contains(out, "($EXAMPLE_FOO_BAR_NAME)")
		r.Contains(out, "-n")
		r.Contains(out, "--name")
	})

	t.Run("check value using flag", func(t *testing.T) {
		cobrautil.SetString(root, "name", "n", "john doe", "name")
		cobrautil.SetString(foo, "name", "n", "john foo", "name")
		cobrautil.SetString(bar, "name", "n", "john bar", "name")

		defer root.ResetFlags()
		defer foo.ResetFlags()
		defer bar.ResetFlags()

		out, err := testutil.CobraExecute(t, root, "")
		r.NoError(err)
		r.Contains(out, "john doe")

		out, err = testutil.CobraExecute(t, root, "foo")
		r.NoError(err)
		r.Contains(out, "john foo")

		out, err = testutil.CobraExecute(t, root, "foo", "bar")
		r.NoError(err)
		r.Contains(out, "john bar")

		out, err = testutil.CobraExecute(t, root, "-n", `"my name is john doe"`)
		r.NoError(err)
		r.Contains(out, "my name is john doe")

		out, err = testutil.CobraExecute(t, root, "foo", "-n", `foo john`)
		r.NoError(err)
		r.Contains(out, "foo john")

		out, err = testutil.CobraExecute(t, root, "foo", "bar", "-n", "bar john")
		r.NoError(err)
		r.Contains(out, "bar john")
	})

	t.Run("check value using flag", func(t *testing.T) {
		cobrautil.SetString(root, "name", "n", "john doe", "name")
		cobrautil.SetString(foo, "name", "n", "john foo", "name")
		cobrautil.SetString(bar, "name", "n", "john bar", "name")

		defer root.ResetFlags()
		defer foo.ResetFlags()
		defer bar.ResetFlags()

		os.Setenv("EXAMPLE_NAME", "my name is john doe")
		os.Setenv("EXAMPLE_FOO_NAME", "foo john")
		os.Setenv("EXAMPLE_FOO_BAR_NAME", "bar john")
		defer os.Clearenv()
		out, err := testutil.CobraExecute(t, root, "")
		r.NoError(err)
		r.Contains(out, "my name is john doe")

		out, err = testutil.CobraExecute(t, root, "foo")
		r.NoError(err)
		r.Contains(out, "foo john")

		out, err = testutil.CobraExecute(t, root, "foo", "bar")
		r.NoError(err)
		r.Contains(out, "bar john")
	})
}
