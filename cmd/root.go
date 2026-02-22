package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var force bool
var rootCmd = &cobra.Command{
	Use:     "snipe [port]",
	Short:   "snipe is a CLI tool for killing processes that are using a specific port.",
	Args:    cobra.ExactArgs(1),
	Example: "snipe 3000\nsnipe 3000 --force",
	RunE: func(cmd *cobra.Command, args []string) error {
		SNIPE(args[0], force)
		return nil
	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force kill the process without confirmation")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing snipe '%s'\n", err)
		os.Exit(1)
	}
}
