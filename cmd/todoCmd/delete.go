package todocmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/favourmusenga/task-manager/internals/todos"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Remove a task using its ID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Implementation for deleting a todo task will go here

		todoId, err := strconv.Atoi(args[0])

		if err != nil {
			return fmt.Errorf("The argument must be an integer")
		}

		if todoId < 1 {
			return fmt.Errorf("Argument must be a number greater than 0")
		}

		dbc, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		ctx := context.Background()

		err = todos.DeleteTodo(dbc, ctx, todoId)

		if err != nil {
			return fmt.Errorf("Category deletion error %w", err)
		}

		fmt.Println("You successful deleted todo")
		return nil
	},
}

func init() {
	todoCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
}
