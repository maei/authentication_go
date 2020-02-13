package models

type User struct {
	Firstname string  `json:"firstName"`
	Lastname  string  `json:"lastName"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Uuid      string  `json:"userId"`
	Address   Address `json:"address"`
}

type Address struct {
	Nr      int    `json:"nr"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Zip     int    `json:"zip"`
	Country string `json:"country"`
}
