package application

import (
	"context"
	"task/src/repository"
)

type UserService interface {
	Create(ctx context.Context, req repository.CreateUserParams) (repository.User, error)
	Update(ctx context.Context, req repository.UpdateUserParams) (int32, error)
	Delete(ctx context.Context, username string) (int32, error)
	GetByUsername(ctx context.Context, username, password string) (string, error)
}
