package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use: "config",
	// Args:  cobra.NoArgs,
	Short: "Get or set config values",
	// Run: func(cmd *cobra.Command, args []string) {
	//     fmt.Println("config called")
	// },
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
}
