package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a task",
	RunE:    deleteRunE,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteRunE(_ *cobra.Command, _ []string) error {
	if err := tg.NewApp(config.Config.DataDir).DeleteTask(filter); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
