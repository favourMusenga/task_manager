package todocmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  "Mark a task as completed using its ID.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Marking a task as completed...")
	},
}

func init() {
	todoCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.
}
