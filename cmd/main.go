package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	database.Init()

	bot, err := tgbotapi.NewBotAPI("6142224756:AAFYO2_mgxeumSMEd6rJjpuhwSufJC7ti7E")
	if err != nil {
		log.Println(err)
	}	

	// bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	telegramBot.Start()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}