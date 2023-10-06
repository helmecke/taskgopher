package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"done"},
	Short:   "Complete a task",
	RunE:    completeRunE,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func completeRunE(_ *cobra.Command, _ []string) error {
	if err := tg.NewApp(config.Config.DataDir).CompleteTask(filter); err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}

	return nil
}
