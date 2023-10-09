package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// config/setCmd represents the config/set command
var configSetCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.ExactArgs(2),
	Short: "Set a config value",
	Run: func(_ *cobra.Command, args []string) {
		viper.Set(args[0], args[1])
		if err := viper.SafeWriteConfig(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Set %s to %s.\n", args[0], args[1])
	},
}
