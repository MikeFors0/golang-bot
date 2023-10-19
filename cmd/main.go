package main

import (
	// "context"
	"log"

	// "github.com/MikeFors0/golang-bot/database"
	// "github.com/MikeFors0/golang-bot/models"
	"github.com/MikeFors0/golang-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	// database.Init()


	// bot.Debug = true
	// user := models.User{
	// 	Login:       "Ira",
	// 	FIO_student: "Smolkin",
	// 	Password:    "1111",
	// }
	
	// user2 := models.User{
	// 	Login:       "Gandon",
	// 	FIO_student: "Smolkin",
	// 	Password:    "2222",
	// }

	// user3 := models.User{
	// 	Login:       "Mike",
	// 	FIO_student: "Smolkin",
	// 	Password:    "3333",
	// }

	// user4 := models.User{
	// 	Login:       "Admin",
	// 	FIO_student: "Smolkin",
	// 	Password:    "4444",
	// }

	// database.AddUser(&user)
	// database.AddUser(&user2)
	// database.AddUser(&user3)
	// database.AddUser(&user4)

	// users, err := database.GetUsers()
	// if err != nil {
	// 	log.Panic(err)
	// 	return
	// }
	// log.Println(users)

	telegram.Init_Context()

	bot, err := tgbotapi.NewBotAPI("6142224756:AAFYO2_mgxeumSMEd6rJjpuhwSufJC7ti7E")
	if err != nil {
		log.Println(err)
	}



	telegramBot := telegram.NewBot(bot)
	telegramBot.Start()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
