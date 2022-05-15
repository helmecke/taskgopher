package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/helmecke/taskgopher/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "taskgopher",
	RunE: list,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err != nil || cmd == nil {
		if _, err := strconv.Atoi(os.Args[1]); err == nil {
			if len(os.Args[1:]) == 1 {
				rootCmd.SetArgs(append([]string{"show"}, os.Args[1:]...))
			}

			if len(os.Args[1:]) == 2 {
				rootCmd.SetArgs([]string{os.Args[2], os.Args[1]})
			}

			if len(os.Args[1:]) > 2 {
				rootCmd.SetArgs(append([]string{os.Args[2], os.Args[1]}, os.Args[3:]...))
			}
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
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
