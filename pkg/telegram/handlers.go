package telegram

import (
	"log"
	"strings"

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



func (b *Bot) handleLogin(message *tgbotapi.Message) (string, string) {
	var (
		login string
		password string
	)


	if text := strings.Split(message.Text, "\n"); len(text) == 2 {
		if err := strings.Split(text[0], "@"); len(err) == 1 {
			Reset_User_Command(message.From.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Логин указан неверно, повторите попытку ещё раз.")
			return "", ""
		}

		login = text[0]
		password = text[1]

	} else if len(text) == 1 {
		_text := strings.Split(message.Text, " ") 
		if len(_text) != 2 {
			Reset_User_Command(message.From.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Данные указаны неверно, повторите попытку ещё раз.")
			return "", ""
		}

		if err := strings.Split(text[0], "@"); len(err) == 1 {
			Reset_User_Command(message.From.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Логин указан неверно, повторите попытку ещё раз.")
			return "", ""
		}

		login = _text[0]
		password = _text[1]
	
	} else { 
		Reset_User_Command(message.From.ID, "reset_login")
		log.Printf("new command: reset_login")
		b.setMessage(message, "Данные указаны неверно, повторите попытку ещё раз.")
		return "", ""
	}

	return login, password
}



