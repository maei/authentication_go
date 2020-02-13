package main

import (
	"log"

	"github.com/maei/authentication_go/router"
)


func main() {
	log.Print("Server Started...")
	e := router.NewRouter()




	e.Start(":8088")
}
