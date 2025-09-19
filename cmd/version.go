package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "dev"
	Commit  = ""
	Date    = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("fsnote version %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
