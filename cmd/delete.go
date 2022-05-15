package cmd

import (
	"fmt"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a task",
	RunE:    delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) error {
	prompt := promptui.Prompt{
		Label:     "Delete task",
		IsConfirm: true,
	}

	if _, err := prompt.Run(); err == nil {
		return tg.NewApp(config.Config.DataDir).Delete(args)
	}

	fmt.Println("Aborted...")
	return nil
}
