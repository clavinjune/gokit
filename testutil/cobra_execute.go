package testutil

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// CobraExecute helps test to execute given cobra.Command
func CobraExecute(t *testing.T, cmd *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return strings.TrimSpace(buf.String()), err
}
