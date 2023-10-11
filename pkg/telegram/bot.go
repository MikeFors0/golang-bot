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

		log.Println(update.Message)

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(int(update.Message.Chat.ID), update.Message)
			continue
		}

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
