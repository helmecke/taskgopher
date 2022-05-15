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
	RunE:    modify,
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}

func modify(cmd *cobra.Command, args []string) error {
	if err := tg.NewApp(config.Config.DataDir).Modify(args); err != nil {
		return fmt.Errorf("failed to modify task: %w", err)
	}

	return nil
}
