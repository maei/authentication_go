package handlers

type User struct{
	Name string `json:"name"`
}

func GetUser() User{
	u1:=User{"Matthias"}
	return u1
}