package controller

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"task/common/helper"
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
	req := new(CreateTaskRequest)

	err := ecx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	task := repository.CreateTaskParams{
		Description: req.Description,
		Status:      sql.NullString{String: *req.Status, Valid: true},
	}
	createTask, err := controller.taskService.Create(ecx.Request().Context(), task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createTaskResponse := helper.Response{
		Code:    http.StatusCreated,
		Message: "OK",
		Data:    createTask,
	}
	return ecx.JSON(http.StatusCreated, createTaskResponse)
}

func (controller *TaskControllerImpl) Update(ecx echo.Context) error {
	id := ecx.Param("id")
	req := new(UpdateTaskRequest)
	sId, _ := strconv.Atoi(id)

	err := ecx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	task := repository.UpdateTaskParams{
		Description: req.Description,
		Status:      sql.NullString{String: *req.Status, Valid: true},
		ID:          int32(sId),
	}
	updateTask, err := controller.taskService.Update(ecx.Request().Context(), task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	updateTaskResponse := helper.Response{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    updateTask,
	}
	return ecx.JSON(http.StatusOK, updateTaskResponse)
}

func (controller *TaskControllerImpl) Delete(ecx echo.Context) error {
	id := ecx.Param("id")
	sId, _ := strconv.Atoi(id)

	err := controller.taskService.Delete(ecx.Request().Context(), int32(sId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	deleteTaskResponse := helper.Response{
		Code:    http.StatusOK,
		Message: "Success Delete Task",
	}
	return ecx.JSON(http.StatusOK, deleteTaskResponse)
}

func (controller *TaskControllerImpl) List(ecx echo.Context) error {
	listTask, err := controller.taskService.List(ecx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	listTaskResponse := helper.Response{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    listTask,
	}
	return ecx.JSON(http.StatusOK, listTaskResponse)
}

func (controller *TaskControllerImpl) ListByStatus(ecx echo.Context) error {
	status := ecx.Param("status")
	newStatus := sql.NullString{String: status, Valid: true}
	listTaskByStatus, err := controller.taskService.ListByStatus(ecx.Request().Context(), newStatus)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	responseListByStatus := helper.Response{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    listTaskByStatus,
	}
	return ecx.JSON(http.StatusOK, responseListByStatus)
}
