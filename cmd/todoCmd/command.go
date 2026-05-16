package todocmd

import (
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todos",
	Short: "Create and manage todo tasks",
	Long:  `Manage your todos from the terminal with support for task tracking, updates, completion status, and priorities.`,
}

func NewCommand() *cobra.Command {
	return todoCmd
}
