package cmd

import (
	"context"
	"fmt"

	"github.com/zeiss/ghmd/internal/cfg"

	"github.com/spf13/cobra"
)

var config = cfg.Default()

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	RootCmd.AddCommand(InitCmd)

	RootCmd.Flags().StringVarP(&config.Root.Run, "run", "r", config.Root.Run, "run a specific hook")

	RootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", config.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&config.Force, "force", "f", config.Force, "force overwrite")
	RootCmd.PersistentFlags().StringVarP(&config.File, "config", "c", config.File, "config file")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "ghmd",
	Short: "ghmd is a markdown generator for GitHub",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(_ context.Context) error {
	return nil // no-op
}
