package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)



//обработчик сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	command_user, err := Get_User_Command(message.From.ID)
	if err != nil {
		return err
	}

	switch command_user.Command {
	case "start":
		return b.Reg(message)
	case "reset_login":
		return b.Reg(message)

	

	default:
		_err := b.setMessage(message, "К сожалению, я не знаю такой команды =(")
		if _err != nil {
			return _err
		}
	}

	return nil
}



//обработчик команд
func (b *Bot) handleCommand(chat_id int, message *tgbotapi.Message) error {

	switch message.Command() {
	case "start":
		return b.handleStart(message)

	case "auth":
		return b.Auth(message)



	default:
		_err := b.setMessage(message, "К сожалению, я не знаю такой команды =(")
		if _err != nil {
			return _err
		}
	}

	return nil
}













func (b *Bot) handleStart(message *tgbotapi.Message) error {

	err := Set_User_Command(int(message.Chat.ID))
	if err != nil {
		return err
	}

	text := "Здравствуй, дорогой пользователь!\nДобро пожаловать в систему помощника по просмотру посещаемости учеников Самарского Государственного Колледжа.\nЯ буду отправлять Вам уведомления, когда Ваш ребёнок придёт в колледж.\nНапишите мне свои логин и пароль как на нашем сайте в любом из форматов ниже:\n\nuser@gmail.com 1234\n\nuser@gmail.com\n1234"

	_err := b.setMessage(message, text)
	if _err != nil {
		return _err
	}

	return nil
}






