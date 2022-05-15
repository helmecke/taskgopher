package cmd

import (
	"fmt"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a task",
	RunE:  stop,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, args []string) error {
	if err := tg.NewApp(config.Config.DataDir).Stop(args); err != nil {
		return fmt.Errorf("failed to stop task: %w", err)
	}

	return nil
}
