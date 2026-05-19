package categories

import (
	"context"

	dbm "github.com/favourmusenga/task-manager/internals/db"
	"gorm.io/gorm"
)

type CategoriesOptions struct {
	Description *string
}

func AddCategory(dbc *gorm.DB, ctx context.Context, name string, categoryOption CategoriesOptions) error {
	err := gorm.G[dbm.Category](dbc).Create(ctx, &dbm.Category{Name: name})

	return err
}
