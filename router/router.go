package router

import (
	"github.com/maei/authentication_go/api/groups"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	// set Groups
	jwtGroups := e.Group("/api/auth")

	groups.PublicGroups(e)
	groups.JwtGroups(jwtGroups)
	
	return e
}