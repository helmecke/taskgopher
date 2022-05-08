package cmd

import (
	"fmt"
	"log"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var due string

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:     "modify",
	Aliases: []string{"mod"},
	Short:   "Modify a task",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlag("due", cmd.Flags().Lookup("due")); err != nil {
			log.Fatal(err)
		}
	},
	RunE: modify,
}

func init() {
	rootCmd.AddCommand(modifyCmd)
	modifyCmd.Flags().StringVarP(&due, "due", "d", "", "due")
}

func modify(cmd *cobra.Command, args []string) error {
	fmt.Println(due)
	return tg.NewTaskgopher(config.Config.DataDir).Modify(args)
}
