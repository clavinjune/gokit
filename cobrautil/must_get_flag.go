package cobrautil

import "github.com/spf13/cobra"

func MustGetBoolFlag(cmd *cobra.Command, name string) bool {
	v, err := cmd.Flags().GetBool(name)
	if err != nil {
		panic(err.Error())
	}

	return v
}

func MustGetIntFlag(cmd *cobra.Command, name string) int {
	v, err := cmd.Flags().GetInt(name)
	if err != nil {
		panic(err.Error())
	}

	return v
}
