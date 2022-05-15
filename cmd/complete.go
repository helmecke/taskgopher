package cmd

import (
	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"done"},
	Short:   "Complete a task",
	Args:    cobra.ExactArgs(1),
	RunE:    complete,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func complete(cmd *cobra.Command, args []string) error {
	return tg.NewApp(config.Config.DataDir).Complete(args)
}
