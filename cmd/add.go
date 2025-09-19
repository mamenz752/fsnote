package cmd

import (
	"fmt"
	"strings"

	"github.com/mamenz752/fsnote/internal/store"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <folder> <memo>",
	Short: "Add or update a memo for a folder",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		folder := args[0]
		memo := strings.Join(args[1:], " ")
		abs, err := store.Add(folder, memo)
		if err != nil {
			return err
		}
		fmt.Printf("Saved memo for %s\n", abs)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
