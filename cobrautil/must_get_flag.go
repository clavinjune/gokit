package cobrautil

import (
	"os"
	"strings"

	"github.com/clavinjune/gokit/argutil"
	"github.com/spf13/cobra"
)

func MustGetBoolFlag(cmd *cobra.Command, name string, checkEnv ...bool) bool {
	setEnvIfNeeded(cmd, name, checkEnv)
	v, err := cmd.Flags().GetBool(name)
	if err != nil {
		panic(err.Error())
	}

	return v
}

func MustGetIntFlag(cmd *cobra.Command, name string, checkEnv ...bool) int {
	setEnvIfNeeded(cmd, name, checkEnv)
	v, err := cmd.Flags().GetInt(name)
	if err != nil {
		panic(err.Error())
	}

	return v
}

func MustGetStringFlag(cmd *cobra.Command, name string, checkEnv ...bool) string {
	setEnvIfNeeded(cmd, name, checkEnv)
	v, err := cmd.Flags().GetString(name)
	if err != nil {
		panic(err.Error())
	}

	return v
}

func setEnvIfNeeded(cmd *cobra.Command, name string, checkEnv []bool) {
	ok := argutil.FirstOrDefault[bool](false, checkEnv...)
	if !ok {
		return
	}

	envName := toEnvName(cmd, name)
	envVal := strings.TrimSpace(os.Getenv(envName))
	if envVal == "" {
		return
	}
	if err := cmd.Flags().Set(name, envVal); err != nil {
		panic(err.Error())
	}
}
