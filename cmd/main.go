package main

import (
	"github.com/maei/authentication_go/router"
)


func main() {
	e := router.NewRouter()
	e.Start(":8088")
}
