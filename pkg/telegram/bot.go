package telegram

import (
	"fmt"
	"log"
	"sync"
	"time"

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

	var wg sync.WaitGroup

	for update := range updates {

		log.Println("\n\nПолучено новое сообщение: " + "\nusername -> " + update.Message.Chat.UserName + "\ntext -> " + update.Message.Text + "\nchat_id -> " + fmt.Sprint(update.Message.Chat.ID) + "\nКонец\n")

		//если обновлений нет, продолжит ожидать
		if update.Message == nil {
			continue
		}
		

		if update.PreCheckoutQuery != nil {
			b.HandlePreCheckoutQuery(&update)
		} else if update.Message != nil && update.Message.SuccessfulPayment != nil {
			b.HandleSuccessfulPayment(update)
		}

		//если это команда, перейдём в обработчик команд
		if update.Message.IsCommand() {
			wg.Add(1)
			go b.handleCommand(update.Message.Chat.ID, update.Message, &wg)
			time.Sleep(time.Second * 1)
			continue
		} else {      //если текст, перейдём в обработчик сообщений
			wg.Add(1)
			go b.handleMessage(update.Message, &wg)
			time.Sleep(time.Second * 1)
			continue
		}

	}

	wg.Wait()
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

// отправить сообщение в чат
// принимает данные чата и текст сообщения, которое мы хотим отправить
func (b *Bot) setMessage(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(int64(message.Chat.ID), text)
	_, err := b.bot.Send(msg)
	return err
}
