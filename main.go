package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
)

func main() {

	user := models.User{
		Login:    "Admin",
		Password: "1234",
	}
	database.AddUser(&user)

	users, err := database.GetUsers()
	if err != nil {
		log.Println(err)
		return
	}
	for _, user := range *users {
		log.Println(user)
	}

}
