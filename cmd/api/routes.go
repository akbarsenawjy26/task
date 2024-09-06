package main

import (
	"github.com/labstack/echo/v4"
	"task/src/task/controller"
)

func RegisterRoutes(ee *echo.Echo, taskController controller.TaskController) {
	api := ee.Group("/api/v1")

	task := api.Group("/task")
	task.POST("/", taskController.Create)
}
