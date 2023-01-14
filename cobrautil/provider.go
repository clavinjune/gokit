package cobrautil

import (
	"io"

	"github.com/clavinjune/gokit/slogutil"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func DefaultPersistentPreRunE(writer io.Writer) func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		logger := slogutil.New(&slogutil.Option{
			IsDebug:      MustGetBoolFlag(cmd, "debug"),
			IsJSON:       MustGetBoolFlag(cmd, "json"),
			SetAsDefault: true,
			Writer:       writer,
		})

		cmd.SetOut(slogutil.NewWriter(logger, slog.LevelInfo))
		cmd.SetErr(slogutil.NewWriter(logger, slog.LevelError))

		return nil
	}
}

func New(name, version string, out io.Writer) *cobra.Command {
	root := &cobra.Command{
		Use:               name,
		Version:           version,
		PersistentPreRunE: DefaultPersistentPreRunE(out),
	}

	root.PersistentFlags().Bool("debug", false, "enable debug mode")
	root.PersistentFlags().Bool("json", false, "enable json mode")
	return root
}
