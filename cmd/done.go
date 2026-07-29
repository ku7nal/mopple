/*
Copyright © 2026 Kunal Kokande
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the task as completed",
Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// get the argument and convert the string to int
		id, _ := strconv.Atoi(args[0])

		// load the file into []tasks
		loadTasks()

		// update the status state of the element in []tasks and write the modified file
		doneTask(id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
