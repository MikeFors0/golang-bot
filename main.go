package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
	// "github.com/MikeFors0/golang-bot/models"
)

func main() {
	user := models.User{
		Login:       "trtr",
		FIO_student: "Smolkin",
		Password:    "1234",
	}
	database.AddUser(&user)
	_, err := database.AddUserTelegram(database.Client, 1469064658)
	if err != nil {
		log.Panic(err)
	}

	_, err = database.AuthenticateUser(database.Client, 1469064658, "trtr", "1234")
	if err != nil {
		log.Panic(err)
	}
	users, err := database.GetUser(1469064658)
	if err != nil {
		log.Panic(err)
		return
	}
	log.Println(users)

	// database.CheckNewData()
	// passage := models.Passage{
	// 	FIO_student: "Smolkin",
	// }
	// time.Sleep(3 * time.Second)
	// res := database.AddData(passage)
	// fmt.Println(res)

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
