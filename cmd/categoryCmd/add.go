package categorycmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a category",
	Long:  "Create a new task category.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding a new category...")
	},
}

func init() {
	categoryCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
}
