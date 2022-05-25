package cmd

import (
	"os"

	"github.com/helmecke/taskgopher/internal/config"
	tg "github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "taskgopher",
	RunE: listRunE,
}

// TODO: replace if cobra 1.5 hits - https://github.com/spf13/cobra/pull/1551
var filter *tg.Filter

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	parser := &tg.Parser{}
	cmd, f, err := parser.ParseArgs(os.Args)
	if err != nil || cmd != "" {
		rootCmd.SetArgs(append([]string{cmd}, os.Args...))
	}

	filter = f

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
