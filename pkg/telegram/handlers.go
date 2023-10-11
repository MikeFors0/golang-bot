package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//обработать сообщение
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(int64(message.From.ID), "Я не знаю такой команды :(")
	_, err := b.bot.Send(msg)
	return err
	
}

func (b *Bot) handleCommand(chat_id int, message *tgbotapi.Message) error {
	user, err := GetUser(chat_id)
	if err != nil {
		return err
	}


	switch user.Runs_Command {
	case "start":
		return b.handleStartCommand(message)

	case "auth":
		return b.handleAuthorizationCommand(message)

	default:
		msg := tgbotapi.NewMessage(int64(message.From.ID), "К сожалению, я не знаю такой команды =(")
		_, err := b.bot.Send(msg)
		return err
	}
}















func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	
	err := SetUser(int(message.Chat.ID))
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(int64(message.From.ID), "Бот запущен :)")
		_, err = b.bot.Send(msg)
		return err
}


func (b *Bot) handleAuthorizationCommand(message *tgbotapi.Message) error {
	Reset_Runs_Command(int(message.Chat.ID), "auth")
	msg := tgbotapi.NewMessage(int64(message.From.ID), "Введите ваши логин и пароль под которвыми вы зарегестрированы в АСУРСО СГК")
		_, err := b.bot.Send(msg)
		return err
}