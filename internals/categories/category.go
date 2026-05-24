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

func ListCategory(dbc *gorm.DB, ctx context.Context) ([]dbm.Category, *gorm.DB) {
	var categories []dbm.Category

	results := dbc.Find(&categories)

	return categories, results
}

func DeleteCategory(dbc *gorm.DB, id int) error {
	results := dbc.Delete(&dbm.Category{}, id)

	return results.Error
}

func UpdateCategory(dbc *gorm.DB, ctx context.Context, id int, name string) error {
	_, err := gorm.G[dbm.Category](dbc).Where("id = ?", id).Update(ctx, "name", name)

	return err
}
