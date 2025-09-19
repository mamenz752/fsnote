package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <folder> <memo>",
	Short: "Add or update a memo for a folder",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("add: (to be implemented)")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
