package telegram

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/MikeFors0/golang-bot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// обработчик сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	var wg sync.WaitGroup

	command_user, err := Get_User_Command(int(message.Chat.ID))
	if err != nil {
		return err
	}

	switch command_user.Command {
	case "start":
		wg.Add(1)
		go b.Reg(message, &wg)

	case "reset_login":
		wg.Add(1)
		go b.Reg(message, &wg)

	default:
		b.setMessage(message, "К сожалению, я не знаю такой команды =(")
	}


	fmt.Scan()

	return nil
}

// обработчик команд
func (b *Bot) handleCommand(chat_id int, message *tgbotapi.Message) error {

	var wg sync.WaitGroup

	switch message.Command() {
	case "start":
		wg.Add(1)
		go b.handleStart(message)

	case "auth":
		wg.Add(1)
		go b.Auth(message, &wg)

	default:
		b.setMessage(message, "К сожалению, я не знаю такой команды =((")
	}

	fmt.Scan()

	return nil
}

func (b *Bot) handleStart(message *tgbotapi.Message) error {

	err := Set_User_Command(int(message.Chat.ID))
	if err != nil {
		return err
	}



	_, err = database.AddUserTelegram(database.Client, uint(message.Chat.ID))
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
		login    string
		password string
	)

	if text := strings.Split(message.Text, "\n"); len(text) == 2 {
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
