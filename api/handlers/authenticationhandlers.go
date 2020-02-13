package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func UserRegistration(c echo.Context) error{
	return c.JSON(http.StatusOK, "Hello World")
}

func UserLogin(c echo.Context) error{
	return c.JSON(http.StatusOK, "Hello World")
}