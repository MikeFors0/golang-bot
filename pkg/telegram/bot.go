package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

// обработчик обновлений
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {

		log.Println("1: " + update.Message.Chat.UserName + " " + "2: " + update.Message.Text)

		//если обновлений нет, продолжит ожидать
		if update.Message == nil {
			continue
		} 

		//если это команда, перейдём в обработчик команд
		if update.Message.IsCommand() {
			b.handleCommand(int(update.Message.From.ID), update.Message)
			continue
		}

		//если текст, перейдём в обработчик сообщений
		b.handleMessage(update.Message)
	}
}

// проверка обновлений (нет ли новых сообщений)
func (b *Bot) initUbdateChanel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	update, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return update, nil
}

//отправить сообщение в чат
//принимает данные чата и текст сообщения, которое мы хотим отправить
func (b *Bot) setMessage(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(int64(message.From.ID), text)
	_, err := b.bot.Send(msg)
	return err
}
