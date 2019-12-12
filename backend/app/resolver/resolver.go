//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"app/config"
	"app/model"
	"context"
)

type Resolver struct{}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func New() *Resolver {
	return &Resolver{}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Tasks(ctx context.Context, input model.TasksInput, orderBy model.TaskOrderFields, page model.PaginationInput) (*model.TaskConnection, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	db := config.DB()

	id, err := config.ShortID().Generate()
	if err != nil {
		return &model.Task{}, err
	}

	task := model.Task{
		Identifier: id,
		Title:      input.Title,
		Due:        input.Due,
	}
	if input.Notes != nil {
		task.Notes = *input.Notes
	}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}
	if err := db.Create(&task).Error; err != nil {
		return &model.Task{}, err
	}

	return &task, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	panic("not implemented")
}
