package cmd

import (
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
	RunE: list,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "A", false, `List all tasks.`)
}

func list(cmd *cobra.Command, args []string) error {
	return tg.NewTaskgopher(config.Config.DataDir).List(viper.GetBool("all"))
}
