package cmd

import (
	"fmt"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a task",
	RunE:  start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) error {
	if err := tg.NewApp(config.Config.DataDir).Start(args); err != nil {
		return fmt.Errorf("failed to start task: %w", err)
	}

	return nil
}
