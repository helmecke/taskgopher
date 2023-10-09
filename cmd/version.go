package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/helmecke/taskgopher/internal/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("version:", version.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
