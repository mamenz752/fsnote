package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsnote",
	Short: "Folder memo management tool on MacOS",
	Long:  "Attach short memos to folders so you remember why they exist.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("fsnote: use a subcommand (add, get, list, remove)")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
