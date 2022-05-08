package cmd

import (
	"fmt"

	"github.com/helmecke/taskgopher/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version:", version.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
