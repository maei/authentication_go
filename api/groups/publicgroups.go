package groups

import (
	"github.com/labstack/echo"
	"github.com/maei/authentication_go/api/handlers"
)

func PublicGroups(c *echo.Echo){
	c.POST("/register", handlers.UserRegistration)
	c.POST("/login", handlers.UserLogin)

}