package main

import (
	"github.com/labstack/echo/v4"
	"task/src/task/controller"
)

func RegisterRoutes(ee *echo.Echo, taskController controller.TaskController) {
	api := ee.Group("/api/v1")

	task := api.Group("/task")
	task.POST("", taskController.Create)
	task.PUT("/:id", taskController.Update)
	task.DELETE("/:id", taskController.Delete)
	task.GET("", taskController.List)
	task.GET("/:status", taskController.ListByStatus)
}
