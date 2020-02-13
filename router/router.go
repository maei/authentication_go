package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/maei/authentication_go/api/groups"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	jwtGroups := e.Group("/api/auth")
	jwtGroups.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("superSecret"),
	}))

	groups.PublicGroups(e)
	groups.JwtGroups(jwtGroups)

	return e
}
