package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/helmecke/taskgopher/internal/config"
	"github.com/helmecke/taskgopher/internal/parser"
)

var (
	cfgFile string
	debug   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:    "taskgopher",
	PreRun: toggleDebug,
	RunE:   listRunE,
}

// TODO: replace if cobra 1.5 hits - https://github.com/spf13/cobra/pull/1551
var (
	filter *parser.Filter
	mod    *parser.Modification
)

type PlainFormatter struct{}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	parser := parser.NewParser()
	err := parser.ParseArgs(os.Args)
	if err != nil || parser.Command != "" {
		// TODO: cleanup args
		rootCmd.SetArgs(append([]string{parser.Command}, os.Args...))
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
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "verbose logging")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.Init(cfgFile)
}

func toggleDebug(cmd *cobra.Command, args []string) {
	if debug {
		log.Info("Debug logs enabled")
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
	} else {
		plainFormatter := new(PlainFormatter)
		log.SetFormatter(plainFormatter)
	}
}
