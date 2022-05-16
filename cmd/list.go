package cmd

import (
	"fmt"
	"log"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlag("all", cmd.Flags().Lookup("all")); err != nil {
			log.Fatal(err)
		}
	},
	RunE: listRunE,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "A", false, `List all tasks.`)
}

func listRunE(_ *cobra.Command, args []string) error {
	if err := tg.NewApp(config.Config.DataDir).ListTasks(args, viper.GetBool("all")); err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}

	return nil
}
