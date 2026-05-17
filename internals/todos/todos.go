package todos

import (
	"context"

	dbm "github.com/favourmusenga/task-manager/internals/db"
	"gorm.io/gorm"
)

type TodosOptions struct {
	dueDate     string
	priotriy    dbm.Priority
	profile     string
	description string
	categories  string
}

func AddTodo(dbc *gorm.DB, ctx context.Context, title string, todoOptions TodosOptions) error {
	err := gorm.G[dbm.Todo](dbc).Create(ctx, &dbm.Todo{Title: title})

	return err
}
