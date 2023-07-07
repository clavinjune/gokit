package cobrautil

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	r = regexp.MustCompile(`\W+`)
)

func prefix(cmd *cobra.Command) string {
	e := cmd.Name()

	for p := cmd.Parent(); p != nil; p = p.Parent() {
		e = fmt.Sprintf("%s_%s", p.Name(), e)
	}

	return strings.ToUpper(
		r.ReplaceAllString(e, "_"),
	)
}

func appendEnvVarToUsage(prefix, name, usage string) string {
	return fmt.Sprintf("%s ($%s_%s)",
		usage, prefix, strings.ToUpper(name),
	)
}

func getEnv(cmd *cobra.Command, name string) string {
	v := fmt.Sprintf("%s_%s", prefix(cmd),
		strings.ToUpper(name),
	)

	return strings.TrimSpace(os.Getenv(v))
}

func SetBool(cmd *cobra.Command, name, shorthand string, value bool, usage string) *bool {
	p := prefix(cmd)
	u := appendEnvVarToUsage(p, name, usage)

	return cmd.Flags().BoolP(name, shorthand, value, u)
}

func MustBool(cmd *cobra.Command, name string) bool {
	e := getEnv(cmd, name)
	if e == "" {
		return must(cmd.Flags().GetBool(name))
	}

	return e == "true"
}

func SetString(cmd *cobra.Command, name, shorthand string, value string, usage string) *string {
	p := prefix(cmd)
	u := appendEnvVarToUsage(p, name, usage)

	return cmd.Flags().StringP(name, shorthand, value, u)
}

func MustString(cmd *cobra.Command, name string) string {
	e := getEnv(cmd, name)
	if e == "" {
		return must(cmd.Flags().GetString(name))
	}

	return e
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}
