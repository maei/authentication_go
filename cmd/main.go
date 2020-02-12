package main

import (
	"fmt"
	"github.com/maei/authentication_go/api/handlers"
)

func main() {
	fmt.Println("Hello World")
	user := handlers.GetUser()
	fmt.Println(user.Name)

}
