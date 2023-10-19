package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
	// "github.com/MikeFors0/golang-bot/models"
)

func main() {
	user := models.User{
		Login:       "Time",
		FIO_student: "gnome",
		Password:    "1234",
	}
	database.AddUser(&user)
	_, err := database.AddUserTelegram(database.Client, 3434)
	if err != nil {
		log.Panic(err)
	}

	_, err = database.AuthenticateUser(database.Client, 3434, "Time", "1234")
	if err != nil {
		log.Panic(err)
	}
	// database.GetUser(1469064658)

	passage := models.Passage{
		FIO_student: "gnome",
	}
	database.AddPassage(passage)
	// database.CheckUserArrayVariable(1212, "[]Passage")
	// database.GetAllPassages()

	database.CheckNewData()
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
