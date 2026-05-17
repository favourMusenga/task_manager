package todocmd

import (
	"context"
	"fmt"

	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/favourmusenga/task-manager/internals/todos"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  "Create a todo task.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		todoOption := todos.TodosOptions{}

		dbc, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		ctx := context.Background()

		err = todos.AddTodo(dbc, ctx, title, todoOption)

		if err != nil {
			return fmt.Errorf("Adding todo error: %w", err)
		}

		fmt.Printf("Added todo successfully to database")

		return nil

	},
}

func init() {
	todoCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
}
