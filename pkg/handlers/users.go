package handlers

import (
	"github.com/elvin-tacirzade/kubernetes-example/pkg/services"
	"github.com/labstack/echo/v4"
)

type Users interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
}

type users struct {
	usersService services.Users
}

func NewUsers(usersService services.Users) Users {
	return &users{
		usersService: usersService,
	}
}

func (u *users) Create(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	email := c.Request().PostFormValue("email")
	response := u.usersService.Create(name, email)
	return c.JSON(response.Code, response.Body)
}

func (u *users) Get(c echo.Context) error {
	response := u.usersService.Get()
	return c.JSON(response.Code, response.Body)
}
