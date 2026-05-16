package todocmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  "Create a todo task.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding a new todo task...")
	},
}

func init() {
	todoCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
}
