package cmd

import (
	"fmt"

	"github.com/mamenz752/fsnote/internal/store"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <folder>",
	Short: "Show memo for a folder",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		abs, err := store.NormalizePath(args[0])
		if err != nil {
			return err
		}
		db, err := store.Load()
		if err != nil {
			return err
		}
		if memo, ok := db[abs]; ok {
			fmt.Printf("%s : %s\n", abs, memo)
		} else {
			fmt.Printf("No memo found for %s", abs)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
