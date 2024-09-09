package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task/app"
	"task/common/helper"
	"task/common/jwt"
	"task/src/repository"
	"task/src/task/application"

	"task/src/task/controller"
	application2 "task/src/users/application"
	controller2 "task/src/users/controller"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := app.NewDB()
	queries := repository.New(database)
	tx := helper.NewTransactionHelper(database)
	taskService := application.NewTaskServiceImpl(queries, tx)
	tokenJWT := jwt.NewJWTServiceImpl()
	userService := application2.NewUserServiceImpl(queries, tx, tokenJWT)
	var taskController controller.TaskController = controller.NewTaskControllerImpl(taskService)
	var userController controller2.UserController = controller2.NewUserControllerImpl(userService)
	RegisterRoutes(e, taskController, userController)

	e.Logger.Fatal(e.Start(":8080"))
}
