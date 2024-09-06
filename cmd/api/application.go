package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task/app"
	"task/common/helper"
	"task/src/repository"
	"task/src/task/application"
	"task/src/task/controller"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := app.NewDB()
	queries := repository.New(database)
	tx := helper.NewTransactionHelper(database)
	taskService := application.NewTaskServiceImpl(queries, tx)
	var taskController controller.TaskController = controller.NewTaskControllerImpl(taskService)
	RegisterRoutes(e, taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
