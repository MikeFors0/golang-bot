package main

import (
	"fmt"

	"github.com/MikeFors0/golang-bot/database"
)

func main() {
	// user := models.User{
	// 	Login:    "Admin",
	// 	Password: "1234",
	// }
	// database.AddUser(&user)
	database.AuthenticateUser("Admin", "1234")

	// users, err := database.GetUsers()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for _, user := range *users {
	// 	log.Println(user)
	// }
	err := database.SendMessage()
	if err != nil {
		fmt.Println("error checking students")
		// handle error
	}
	fmt.Println("Its Okey")

}
