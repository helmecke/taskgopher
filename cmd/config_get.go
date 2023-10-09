package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// config/setCmd represents the config/set command
var configGetCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.ExactArgs(1),
	Short: "Get a config value",
	Run: func(_ *cobra.Command, args []string) {
		fmt.Println(viper.Get(args[0]))
	},
}
