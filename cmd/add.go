/*
Copyright © 2026 Kunal Kokande
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the task",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		// get the argument string
		description := args[0]

		// get the csv file
		getFile()

		// load the file into []tasks	
		loadTasks()

		// add the task into the []tasks and append it to csv file
		moppleCreate(description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
