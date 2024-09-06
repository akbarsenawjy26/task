package controller

import (
	"github.com/labstack/echo/v4"
	"task/common/httpservice"
	"task/common/utility"
	"task/src/repository"
	"task/src/task/application"
)

type TaskControllerImpl struct {
	taskService application.TaskService
}

func NewTaskControllerImpl(taskService application.TaskService) TaskController {
	return &TaskControllerImpl{
		taskService: taskService,
	}
}

func (controller *TaskControllerImpl) Create(ecx echo.Context) error {
	var req repository.CreateTaskParams
	if err := ecx.Bind(&req); err != nil {
		return httpservice.ResponseData(ecx, nil, httpservice.ErrBadRequest)
	}
	if err := utility.ValidateStruct(ecx.Request().Context(), req); err != nil {
		return httpservice.ResponseData(ecx, nil, httpservice.ErrBadRequest)
	}

	task, err := controller.taskService.Create(ecx.Request().Context(), req)
	if err != nil {
		return httpservice.ResponseData(ecx, nil, httpservice.ErrInternalServerError)
	}

	return httpservice.ResponseData(ecx, task, nil)
}

func (controller *TaskControllerImpl) Update(ecx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *TaskControllerImpl) Delete(ecx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *TaskControllerImpl) List(ecx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *TaskControllerImpl) ListByStatus(ecx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
