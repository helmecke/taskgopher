package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/helmecke/taskgopher/internal/config"
	"github.com/helmecke/taskgopher/internal/parser"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "taskgopher",
	RunE: listRunE,
}

// TODO: replace if cobra 1.5 hits - https://github.com/spf13/cobra/pull/1551
var (
	filter *parser.Filter
	mod    *parser.Modification
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	parser := parser.NewParser()
	err := parser.ParseArgs(os.Args[1:])
	if err != nil || parser.Command != "" {
		// TODO: cleanup args
		rootCmd.SetArgs([]string{parser.Command})
	}

	filter = parser.Filter
	mod = parser.Modification

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskgopher.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.Init(cfgFile)
}
