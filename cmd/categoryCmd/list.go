package categorycmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/favourmusenga/task-manager/internals/categories"
	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List categories",
	Long:  "View all task categories.",
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := db.GetDb("./task_manager.db")

		if err != nil {
			return fmt.Errorf("Database connection error: %w", err)
		}

		ctx := context.Background()

		categoriesArr, results := categories.ListCategory(db, ctx)

		if results.Error != nil {
			return fmt.Errorf("Failed to fetch categories")
		}

		fmt.Printf("\n%-4s %-30s %s\n", "ID", "NAME", "CREATED")

		fmt.Println(strings.Repeat("-", 90))

		for _, category := range categoriesArr {
			fmt.Printf("%-4d %-30s %s\n", category.ID, category.Name, category.CreatedAt.Format("02-Jan-2006"))
		}

		return nil

	},
}

func init() {
	categoryCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
}
