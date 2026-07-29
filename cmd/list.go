/*
Copyright © 2026 Kunal Kokande

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {

		// load file into []tasks
		loadTasks()

		// read the []tasks and format it into table
		moppleList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
