package todocmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/favourmusenga/task-manager/internals/todos"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  "Mark a task as completed using its ID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Marking a task as completed...")

		todoId, err := strconv.Atoi(args[0])

		if err != nil {
			return fmt.Errorf("The argment must be an integer")
		}

		if todoId < 1 {
			return fmt.Errorf("Argment must be a number greater than 0")
		}

		dbc, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		ctx := context.Background()

		err = todos.CompleteTodo(dbc, ctx, todoId)

		if err != nil {
			return fmt.Errorf("Category deletion error %w", err)
		}

		fmt.Println("You successful Completed todo")
		return nil
	},
}

func init() {
	todoCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.
}
