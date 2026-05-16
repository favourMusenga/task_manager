package todocmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show task summary",
	Long:  "View todo statistics and summaries.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating a summary of your tasks...")
	},
}

func init() {
	todoCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.
}
