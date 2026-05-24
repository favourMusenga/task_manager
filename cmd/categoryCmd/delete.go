package categorycmd

import (
	"fmt"
	"strconv"

	"github.com/favourmusenga/task-manager/internals/categories"
	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a category",
	Long:  "Remove a category.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deleting a category...")

		categoryId, err := strconv.Atoi(args[0])

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

		err = categories.DeleteCategory(dbc, categoryId)

		if err != nil {
			return fmt.Errorf("Category deletion error %w", err)
		}

		fmt.Println("You successfull deleted category")
		return nil

	},
}

func init() {
	categoryCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
}
