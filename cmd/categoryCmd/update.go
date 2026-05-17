package categorycmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a category",
	Long:  "Modify an existing category.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating category...")
	},
}

func init() {
	categoryCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
}
