package todocmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/favourmusenga/task-manager/internals/todos"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  "Modify an existing task.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Updating task...")

		todoId, err := strconv.Atoi(args[0])

		todoOption := todos.TodosOptions{}

		if err != nil {
			return fmt.Errorf("The argument must be an integer")
		}

		if todoId < 1 {
			return fmt.Errorf("Argument must be a number greater than 0")
		}

		todoName, err := cmd.Flags().GetString("name")

		if err != nil {
			return err
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

		dbc, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		priority = strings.ToLower(priority)
		profile = strings.ToLower(profile)

		if priority != "" {
			if priority != "low" && priority != "medium" && priority != "high" {
				return fmt.Errorf("Priority must be 'low', 'medium' or 'high'")
			}
			todoOption.Priority = db.Priority(priority)

		}

		if profile != "" {
			if profile != "work" && profile != "personal" {
				return fmt.Errorf("Profile must be 'work' or 'personal'")
			}
			todoOption.Profile = db.Profile(profile)
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

		err = todos.UpdateTodo(dbc, todoId, todoName, todoOption)

		if err != nil {
			return fmt.Errorf("Updating todo error: %w", err)
		}

		fmt.Printf("Updated todo successfully")

		return nil

	},
}

func init() {
	todoCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
	updateCmd.Flags().UintP("category-id", "c", 0, "Add Category to task")
	updateCmd.Flags().StringP("due-date", "d", "", "Add due date in YYYY-MM-DD format like 2026-02-01")
	updateCmd.Flags().StringP("priority", "p", "", "Add Priority")
	updateCmd.Flags().StringP("name", "n", "", "update name")
	updateCmd.Flags().String("profile", "", "Choose profile")
}
