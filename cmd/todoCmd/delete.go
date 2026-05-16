package todocmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Remove a task using its ID.",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for deleting a todo task will go here
	},
}

func init() {
	todoCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
}
