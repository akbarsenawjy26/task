package controller

import "github.com/labstack/echo/v4"

type TaskController interface {
	Create(ecx echo.Context) error
	Update(ecx echo.Context) error
	Delete(ecx echo.Context) error
	List(ecx echo.Context) error
	ListByStatus(ecx echo.Context) error
}
