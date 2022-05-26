package cmd

import (
	"fmt"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:     "modify",
	Aliases: []string{"mod"},
	Short:   "Modify a task",
	RunE:    modifyRunE,
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}

func modifyRunE(_ *cobra.Command, _ []string) error {
	if err := tg.NewApp(config.Config.DataDir).ModifyTask(filter, mod); err != nil {
		return fmt.Errorf("failed to modify task: %w", err)
	}

	return nil
}
