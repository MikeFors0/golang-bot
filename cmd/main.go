package main

import (
	"log"
	"os"

	"github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/models"
	"github.com/MikeFors0/golang-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("This is a log message")

	telegram.User_comand = make(map[int64]string)
	database.AddPassage(models.Passage{FIO_student: "Smolkin"})
	// user := models.User{
	// 	Login:       "Ira",
	// 	FIO_student: "Smolkin",
	// 	Password:    "1111",
	// }

	user2 := models.User{
		Login:       "Gandon",
		FIO_student: "Smolkin",
		Password:    "2222",
	}

	// user3 := models.User{
	// 	Login:       "Mike",
	// 	FIO_student: "Smolkin",
	// 	Password:    "3333",
	// }

	user4 := models.User{
		Login:       "Admin",
		FIO_student: "Youre",
		Password:    "4444",
	}

	// database.AddUser(&user)
	database.AddUser(&user2)
	// database.AddUser(&user3)
	database.AddUser(&user4)

	bot, err := tgbotapi.NewBotAPI("6142224756:AAFYO2_mgxeumSMEd6rJjpuhwSufJC7ti7E")
	if err != nil {
		log.Println(err)
	}

	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

	telegramBot.GetPassage()

}
