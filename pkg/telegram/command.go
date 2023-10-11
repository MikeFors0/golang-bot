package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) Start() error {
	log.Printf("Authorized on account %v", &b.bot.Self.UserName)

	update, err := b.initUbdateChanel()
	if err != nil {
		return err
	}

	b.handleUpdates(update)

	return nil
}

func (b *Bot) Auth(message *tgbotapi.Message, login, password string) error {

	if text := strings.Split(message.Text, "\n"); len(text) == 2 {
		login = text[0]
		password = text[1]
	}

	if text := strings.Split(message.Text, " "); len(text) == 2 {
		login = text[0]
		password = text[1]
	}

	msg := tgbotapi.NewMessage(int64(message.From.ID), "Ваш логин:"+login+"\n"+"Ваш проль"+password)
	_, err := b.bot.Send(msg)
	return err
}
