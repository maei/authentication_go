package models

import (
	"encoding/json"
	"log"

	"github.com/satori/go.uuid"
)

type User struct {
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Uuid      uuid.UUID `json:"userId"`
	Address   Address   `json:"address"`
}

type Address struct {
	Nr      int    `json:"nr"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Zip     int    `json:"zip"`
	Country string `json:"country"`
}

func (u *User) CreateUser() {
	bs, err := json.Marshal(u)
	if err != nil {
		log.Print(err)
	}
	log.Printf(string(bs))
}
