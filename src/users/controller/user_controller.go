package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Create(ecx echo.Context) error
	Update(ecx echo.Context) error
	Delete(ecx echo.Context) error
	GetByUsername(ecx echo.Context) error
}
