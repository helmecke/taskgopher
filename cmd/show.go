package cmd

import (
	"fmt"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details of a task",
	RunE:  showRunE,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func showRunE(_ *cobra.Command, args []string) error {
	if err := tg.NewApp(config.Config.DataDir).ShowTask(args); err != nil {
		return fmt.Errorf("failed to show task: %w", err)
	}

	return nil
}
