package cmd

import (
	"fmt"

	"github.com/mamenz752/fsnote/internal/store"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <folder>",
	Short: "Remove the memo for a folder",
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
		if _, ok := db[abs]; ok {
			delete(db, abs)
			if err := store.Save(db); err != nil {
				return err
			}
			fmt.Printf("Removed memo for %s\n", abs)
			return nil
		} else {
			fmt.Printf("No memo found for %s\n", abs)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
