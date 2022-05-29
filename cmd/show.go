package cmd

import (
	"errors"
	"fmt"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

var errShowIDs = errors.New("requires a filter that returns one task")

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details of a task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(filter.IDs)+len(filter.UUIDs) != 1 {
			return errShowIDs
		}

		return nil
	},
	RunE: showRunE,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func showRunE(_ *cobra.Command, _ []string) error {
	if err := tg.NewApp(config.Config.DataDir).ShowTask(filter); err != nil {
		return fmt.Errorf("failed to show task: %w", err)
	}

	return nil
}
