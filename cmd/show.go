package cmd

import (
	"fmt"
	"io/ioutil"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details of a task",
	Args:  cobra.ExactArgs(1),
	RunE:  show,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func show(cmd *cobra.Command, args []string) error {
	path := "README.md"
	source, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result := markdown.Render(string(source), 80, 6)

	fmt.Printf("%s", result)

	return nil
}
