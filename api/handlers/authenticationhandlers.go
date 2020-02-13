package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
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

	var wg sync.WaitGroup
	wg.Add(2)

	// cast empty User
	u := models.ToUser(bs)

	// create uuid
	go func(u *models.User, wg *sync.WaitGroup) {
		defer wg.Done()
		uuid := uuid.NewV4()
		(*u).Uuid = uuid.String()
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
	
	// cast empty User
	u := models.ToUser(bs)

	if u.Password != "urlaub" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "not authorized",
		})
	}

	// create a JWT with uuid and JWT-Standardclaims
	access_token, err := models.CreateToken(u.Uuid)
	if err != nil {
		log.Print(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":         u,
		"access_token": access_token,
	})
}
