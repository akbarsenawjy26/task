package controller

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"task/common/helper"
	"task/src/repository"
	"task/src/users/application"
)

type UserControllerImpl struct {
	userService application.UserService
}

func NewUserControllerImpl(userService application.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) Create(ecx echo.Context) error {
	email := ecx.FormValue("email")
	username := ecx.FormValue("username")
	password := ecx.FormValue("password")

	if email == "" || username == "" || password == "" {
		return ecx.JSON(http.StatusUnprocessableEntity, "Email or Username or Password is empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ecx.JSON(http.StatusInternalServerError, err.Error())
	}

	user := repository.CreateUserParams{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}
	createUser, err := controller.userService.Create(ecx.Request().Context(), user)
	if err != nil {
		return ecx.JSON(http.StatusBadRequest, err.Error())
	}
	createUserResponse := helper.Response{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    createUser,
	}
	return ecx.JSON(http.StatusCreated, createUserResponse)
}

func (controller *UserControllerImpl) Update(ecx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Delete(ecx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) GetByUsername(ecx echo.Context) error {
	username := ecx.FormValue("username")
	password := ecx.FormValue("password")

	if username == "" || password == "" {
		return ecx.JSON(http.StatusUnprocessableEntity, "Username or Password is empty")
	}

	getUserByUsername, err := controller.userService.GetByUsername(ecx.Request().Context(), username, password)
	if err != nil {
		return ecx.JSON(http.StatusBadRequest, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = getUserByUsername
	cookie.HttpOnly = true
	ecx.SetCookie(cookie)

	getUserByUsernameResponse := helper.Response{
		Code:    http.StatusAccepted,
		Message: "Success Login",
		Data:    getUserByUsername,
	}
	return ecx.JSON(http.StatusAccepted, getUserByUsernameResponse)
}
