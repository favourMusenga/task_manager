package categorycmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a category",
	Long:  "Remove a category.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting a category...")
	},
}

func init() {
	categoryCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
}
