package router

import (
	"github.com/maei/authentication_go/api/groups"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	jwtGroups := e.Group("/api/auth")

	groups.PublicGroups(e)
	groups.JwtGroups(jwtGroups)

	return e
}