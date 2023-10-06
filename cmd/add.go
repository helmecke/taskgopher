package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	RunE:  addRunE,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addRunE(_ *cobra.Command, _ []string) error {
	if err := tg.NewApp(config.Config.DataDir).AddTask(mod); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	return nil
}
