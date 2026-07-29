/*
Copyright © 2026 Kunal Kokande
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the task",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// get the areguement string and convert it to int 
		taskid , _:= strconv.Atoi(args[0])

		// load the file into []tasks
		loadTasks()

		// delete the element/task from []tasks and replace with modified file
		moppleDelete(taskid)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
