package todocmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  "Modify an existing task.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating task...")
	},
}

func init() {
	todoCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
}
