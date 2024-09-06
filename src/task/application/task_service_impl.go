package application

import (
	"context"
	"database/sql"
	"task/common/helper"
	"task/src/repository"
)

type TaskServiceImpl struct {
	taskRepository    *repository.Queries
	transactionHelper *helper.TransactionHelper
}

func NewTaskServiceImpl(taskRepository *repository.Queries, transactionHelper *helper.TransactionHelper) TaskService {
	return &TaskServiceImpl{
		taskRepository:    taskRepository,
		transactionHelper: transactionHelper,
	}
}

func (service *TaskServiceImpl) Create(ctx context.Context, req repository.CreateTaskParams) (repository.Task, error) {
	var task repository.Task
	err := service.transactionHelper.WithTransaction(ctx, func(tx *sql.Tx) error {
		var err error
		task, err = service.taskRepository.CreateTask(ctx, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return repository.Task{}, err
	}
	return task, nil
}

func (service *TaskServiceImpl) Update(ctx context.Context, req repository.UpdateTaskParams) (int32, error) {
	var result int32
	err := service.transactionHelper.WithTransaction(ctx, func(tx *sql.Tx) error {
		var err error
		result, err = service.taskRepository.UpdateTask(ctx, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (service *TaskServiceImpl) Delete(ctx context.Context, id int32) error {
	return service.transactionHelper.WithTransaction(ctx, func(tx *sql.Tx) error {
		return service.taskRepository.DeleteTask(ctx, id)
	})
}

func (service *TaskServiceImpl) List(ctx context.Context) ([]repository.Task, error) {
	return service.taskRepository.ListTask(ctx)
}

func (service *TaskServiceImpl) ListByStatus(ctx context.Context, status sql.NullString) ([]repository.Task, error) {
	return service.taskRepository.ListTaskByStatus(ctx, status)
}
