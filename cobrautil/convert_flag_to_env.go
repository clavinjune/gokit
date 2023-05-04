package cobrautil

import (
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

const (
	reReplaceSymbolPattern string = `\W+`
)

var (
	reReplaceSymbol = regexp.MustCompile(reReplaceSymbolPattern)
)

func toEnvName(cmd *cobra.Command, name string) string {
	n := reReplaceSymbol.ReplaceAllString(name, "_")

	if cmd.Name() != "" {
		n = cmd.Name() + "_" + n
	}

	return strings.ToUpper(n)
}
