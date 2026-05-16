package categorycmd

import (
	"github.com/spf13/cobra"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Organize tasks into categories",
	Long:  `Create and manage categories to better organize and group your todo tasks.`,
}

func NewCommand() *cobra.Command {
	return categoryCmd
}
