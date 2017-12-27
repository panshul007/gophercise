package main

import (
	"fmt"
	"gophercise/lenslocked.com/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@email.io",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	user.Email = "michael@paperco.io"
	if err := us.Update(&user); err != nil {
		panic(err)
	}
	userById, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userById)
}
