package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/maei/authentication_go/models"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistration(c echo.Context) error {
	defer c.Request().Body.Close()
	bs, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Print(err)
	}

	u := models.User{}

	var wg sync.WaitGroup
	wg.Add(2)

	err = json.Unmarshal(bs, &u)
	if err != nil {
		log.Print(err)
	}

	// create uuid
	go func(u *models.User, wg *sync.WaitGroup) {
		defer wg.Done()
		(*u).Uuid = uuid.NewV4()
	}(&u, &wg)

	// encrypt password
	go func(u *models.User, wg *sync.WaitGroup) {
		defer wg.Done()
		pass, err := bcrypt.GenerateFromPassword([]byte((*u).Password), bcrypt.DefaultCost)
		if err != nil {
			log.Print(err)
		}
		(*u).Password = string(pass)
	}(&u, &wg)

	wg.Wait()

	// Write User to DB
	go u.CreateUser()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user created",
	})
}

func UserLogin(c echo.Context) error {
	defer c.Request().Body.Close()

	bs, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Print(err)
	}

	u := models.User{}
	err = json.Unmarshal(bs, &u)
	if err != nil {
		log.Print(err)
	}

	if u.Password != "urlaub" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "not authorized",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  u,
		"token": "token",
	})
}
