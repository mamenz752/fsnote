package cmd

import (
	"fmt"
	"sort"

	"github.com/mamenz752/fsnote/internal/store"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all folder memos",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := store.Load()
		if err != nil {
			return err
		}
		if len(db) == 0 {
			fmt.Println("No memos yet.")
			return nil
		}
		keys := make([]string, 0, len(db))
		for k := range db {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%s : %s\n", k, db[k])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
