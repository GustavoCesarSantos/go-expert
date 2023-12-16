package handlers

import (
	"net/http"

	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/dtos"
	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/models"
	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/repositories"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	userDTO := dtos.CreateUserRequest{}
	err := c.Bind(&userDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user := models.NewUser(
		userDTO.Name,
		userDTO.Email,
		userDTO.Password,
	)
	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}