package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct{
	Name string `json:"name"`
}

func GetUser(c echo.Context) error{
	// u1:=User{"Matthias"}
	// return u1
	return c.JSON(http.StatusOK, "Hello World")

}

func UpdateUser(c echo.Context) error{
	// u1:=User{"Matthias"}
	// return u1
	return c.JSON(http.StatusOK, "Hello World")

}

func DeleteUser(c echo.Context) error{
	// u1:=User{"Matthias"}
	// return u1
	return c.JSON(http.StatusOK, "Hello World")

}

func Users(c echo.Context) error{
	// u1:=User{"Matthias"}
	// return u1
	return c.JSON(http.StatusOK, "Hello World")

}