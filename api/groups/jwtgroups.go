package groups

import (
	"github.com/labstack/echo"
	"github.com/maei/authentication_go/api/handlers"
)

func JwtGroups(g *echo.Group) {
	g.GET("/user", handlers.GetUser)
	g.PUT("/user", handlers.GetUser)
	g.DELETE("/user", handlers.GetUser)

	g.GET("/users", handlers.Users)

}