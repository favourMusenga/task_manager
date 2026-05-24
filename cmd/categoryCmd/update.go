package categorycmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/favourmusenga/task-manager/internals/categories"
	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a category",
	Long:  "Modify an existing category.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Updating category...")

		categoryId, err := strconv.Atoi(args[0])
		name := args[1]

		if err != nil {
			return fmt.Errorf("The argment must be an integer")
		}

		if categoryId < 1 {
			return fmt.Errorf("Argment must be a number greater than 0")
		}

		dbc, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		ctx := context.Background()

		err = categories.UpdateCategory(dbc, ctx, categoryId, name)

		if err != nil {
			return fmt.Errorf("Category update error %w", err)
		}

		fmt.Println("Category successfully updated")

		return nil

	},
}

func init() {
	categoryCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
}
