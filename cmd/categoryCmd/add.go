package categorycmd

import (
	"context"
	"fmt"

	"github.com/favourmusenga/task-manager/internals/categories"
	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a category",
	Long:  "Create a new task category.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Adding a new category...")

		categoryName := args[0]

		db, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		ctx := context.Background()

		categoryOptions := categories.CategoriesOptions{}

		categoryDescription, err := cmd.Flags().GetString("description")

		if err != nil {
			return err
		}

		if categoryDescription != "" {
			categoryOptions.Description = &categoryDescription
		}

		err = categories.AddCategory(db, ctx, categoryName, categoryOptions)

		if err != nil {
			return fmt.Errorf("Failed to add categories %w", err)
		}

		fmt.Printf("Added category to database")

		return nil

	},
	Args: cobra.ExactArgs(1),
}

func init() {
	categoryCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringP("description", "d", "", "Add categoory description")
}
