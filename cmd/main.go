package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/models"
	"github.com/MikeFors0/golang-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	telegram.User_comand = make(map[int64]string)
	// database.AddPassage(models.Passage{FIO_student: "Smolkin"})

	user := models.User{
		Login:       "Ira",
		FIO_student: "Smolkin",
		Password:    "1111",
	}

	user2 := models.User{
		Login:       "Gandon",
		FIO_student: "GGG",
		Password:    "2222",
	}

	user3 := models.User{
		Login:       "Mike",
		FIO_student: "MMM",
		Password:    "3333",
	}

	user4 := models.User{
		Login:       "Admin",
		FIO_student: "Admin",
		Password:    "4444",
	}

	database.AddUser(&user)
	database.AddUser(&user2)
	database.AddUser(&user3)
	database.AddUser(&user4)

	bot, err := tgbotapi.NewBotAPI("6142224756:AAFYO2_mgxeumSMEd6rJjpuhwSufJC7ti7E")
	if err != nil {
		log.Println(err)
	}

	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Println(err)
	}

	// telegramBot.GetPassage()

	// database.AddPassage(models.Passage{FIO_student: "Smolkin"})

	// time.Sleep(time.Second * 7)

	// database.AddPassage(models.Passage{FIO_student: "Smolkin"})

	// time.Sleep(time.Second * 7)

	// database.AddPassage(models.Passage{FIO_student: "Smolkin"})
}
