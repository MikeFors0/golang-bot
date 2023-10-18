package main

import (
	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
	// "github.com/MikeFors0/golang-bot/models"
)

func main() {
	user := models.User{
		Login:       "Rery",
		FIO_student: "Rerer",
		Password:    "1234",
	}
	database.AddUser(&user)
	// _, err := database.AddUserTelegram(database.Client, 1469064658)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// _, err = database.AuthenticateUser(database.Client, 1469064658, "Admin", "1234")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// database.GetUser(1469064658)

	passage := models.Passage{
		FIO_student: "Rerer",
	}
	database.AddPassage(passage)

	// database.GetAllPassages()

	// database.CheckNewData()
	// database.GetUserByFIO("Smolkin")
	// time.Sleep(3 * time.Second)
	// rr := database.AddData(passage)
	// fmt.Println(rr)
	// for _, user := range users {
	// 	log.Println(user)
	// }
	// err := database.SendMessage()
	// if err != nil {
	// 	fmt.Println("error checking students")
	// 	// handle error
	// }
	// fmt.Println("Its Okey")

}
