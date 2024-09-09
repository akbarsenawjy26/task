package main

import (
	"github.com/labstack/echo/v4"
	"task/common/jwt"
	md "task/common/middleware"
	"task/src/task/controller"
	controller2 "task/src/users/controller"
)

func RegisterRoutes(ee *echo.Echo, taskController controller.TaskController, userController controller2.UserController) {
	jwtService := jwt.NewJWTServiceImpl()

	api := ee.Group("/api/v1")

	task := api.Group("/task", md.JWTMiddleware(jwtService))
	task.POST("", taskController.Create)
	task.PUT("/:id", taskController.Update)
	task.DELETE("/:id", taskController.Delete)
	task.GET("", taskController.List)
	task.GET("/:status", taskController.ListByStatus)

	user := api.Group("/user")
	user.POST("/register", userController.Create)
	user.POST("/login", userController.GetByUsername)
}
