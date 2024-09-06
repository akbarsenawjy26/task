package application

import (
	"context"
	"database/sql"
	"task/src/repository"
)

type TaskService interface {
	Create(ctx context.Context, req repository.CreateTaskParams) (repository.Task, error)
	Update(ctx context.Context, req repository.UpdateTaskParams) (int32, error)
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context) ([]repository.Task, error)
	ListByStatus(ctx context.Context, status sql.NullString) ([]repository.Task, error)
}
