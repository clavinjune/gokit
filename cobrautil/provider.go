package cobrautil

import (
	"github.com/clavinjune/gokit/argutil"
	"github.com/clavinjune/gokit/slogutil"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func DefaultPersistentPreRunE(opt Option) func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		slogOpt := opt.SlogOption
		slogOpt.IsDebug = MustGetBoolFlag(cmd, "debug", true)
		slogOpt.IsJSON = MustGetBoolFlag(cmd, "json", true)
		logger := slogutil.New(&slogOpt)
		slogutil.Put(cmd.Context(), logger)

		if opt.SetOutToSlog {
			cmd.SetOut(slogutil.NewWriter(cmd.Context(), logger, slog.LevelInfo))
		}
		if opt.SetErrToSlog {
			cmd.SetErr(slogutil.NewWriter(cmd.Context(), logger, slog.LevelError))
		}

		return nil
	}
}

func New(name, version string, opts ...*Option) *cobra.Command {
	opt := argutil.FirstOrDefault(&DefaultOption, opts...)

	root := &cobra.Command{
		Use:               name,
		Version:           version,
		PersistentPreRunE: DefaultPersistentPreRunE(*opt),
	}

	root.PersistentFlags().Bool("debug", false, "enable debug mode")
	root.PersistentFlags().Bool("json", false, "enable json mode")
	return root
}
