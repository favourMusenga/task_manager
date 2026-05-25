package todos

import (
	"context"
	"time"

	dbm "github.com/favourmusenga/task-manager/internals/db"
	"gorm.io/gorm"
)

type TodosOptions struct {
	DueDate     time.Time
	Priority    dbm.Priority
	Profile     dbm.Profile
	Description string
	CategoryId  uint
}

func AddTodo(dbc *gorm.DB, ctx context.Context, title string, todoOptions TodosOptions) error {
	var categoryId *uint
	var dueDate *time.Time

	if todoOptions.CategoryId > 0 {
		categoryId = &todoOptions.CategoryId
	}

	if !todoOptions.DueDate.IsZero() {
		dueDate = &todoOptions.DueDate
	}

	err := gorm.G[dbm.Todo](dbc).Create(ctx, &dbm.Todo{Title: title, Priority: todoOptions.Priority, Profile: todoOptions.Profile, DueDate: dueDate, CategoryID: categoryId})

	return err
}

func UpdateTodo(dbc *gorm.DB, id int, title string, todoOptions TodosOptions) error {
	var todo dbm.Todo
	res := dbc.First(&todo, id)

	if res.Error != nil {
		return res.Error
	}

	var categoryId *uint
	var dueDate *time.Time

	if todoOptions.CategoryId > 0 {
		categoryId = &todoOptions.CategoryId
	}

	if !todoOptions.DueDate.IsZero() {
		dueDate = &todoOptions.DueDate
	}

	res = dbc.Model(&todo).Updates(dbm.Todo{Title: title, Priority: todoOptions.Priority, Profile: todoOptions.Profile, DueDate: dueDate, CategoryID: categoryId})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func CompleteTodo(dbc *gorm.DB, ctx context.Context, id int) error {
	_, err := gorm.G[dbm.Todo](dbc).Where("id = ? AND completed =?", id, false).Update(ctx, "completed", true)
	return err
}

func DeleteTodo(dbc *gorm.DB, ctx context.Context, id int) error {
	_, err := gorm.G[dbm.Todo](dbc).Where("id = ?", id).Delete(ctx)

	return err
}
