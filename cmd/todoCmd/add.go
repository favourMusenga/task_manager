package todocmd

import (
	"context"
	"fmt"
	"strings"
	"time"

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

		priority, err := cmd.Flags().GetString("priority")

		if err != nil {
			return err
		}

		profile, err := cmd.Flags().GetString("profile")

		if err != nil {
			return err
		}

		categoryId, err := cmd.Flags().GetUint("category-id")

		if err != nil {
			return err
		}

		dueDate, err := cmd.Flags().GetString("due-date")

		if err != nil {
			return err
		}

		priority = strings.ToLower(priority)
		profile = strings.ToLower(profile)

		if priority != "low" && priority != "medium" && priority != "high" {
			return fmt.Errorf("Priority must be 'low', 'medium' or 'high'")
		}

		if profile != "work" && profile != "personal" {
			return fmt.Errorf("Profile must be 'work' or 'personal'")
		}

		if categoryId > 0 {
			todoOption.CategoryId = categoryId
		}

		if dueDate != "" {
			t, err := time.Parse("2006-01-02", dueDate)

			if err != nil {
				return nil
			}

			todoOption.DueDate = t
		}

		todoOption.Priority = db.Priority(priority)
		todoOption.Profile = db.Profile(profile)

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
	addCmd.Flags().UintP("category-id", "c", 0, "Add Category to task")
	addCmd.Flags().StringP("due-date", "d", "", "Add due date in YYYY-MM-DD format like 2026-02-01")
	addCmd.Flags().StringP("priority", "p", "low", "Add Priority")
	addCmd.Flags().String("profile", "personal", "Choose profile")

}
