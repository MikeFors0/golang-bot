package telegram

import (
	"log"
	"strings"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//запускает бота
func (b *Bot) Start() error {
	log.Printf("Authorized on account %v", &b.bot.Self.UserName)

	update, err := b.initUbdateChanel()
	if err != nil {
		return err
	}

	b.handleUpdates(update)

	return nil
}

//переделать под работу с БД!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//авторизация, пользователь вводит свой логин и пароль 
//может сделать это через пробел, или с переносом на следующую строку
func (b *Bot) Reg(message *tgbotapi.Message) error {
	var (
		login string
		password string
	)

	if text := strings.Split(message.Text, "\n"); len(text) == 2 {
		if err := strings.Split(text[0], "@"); len(err) == 1 {
			Reset_User_Command(message.From.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Логин указан неверно, пожалуйста, повторите попытку.\nНапишите логин и пароль ещё раз.")
			return nil
		}

		login = text[0]
		password = text[1]

	} else if len(text) == 1 {
		_text := strings.Split(message.Text, " ") 
		if len(_text) != 2 {
			Reset_User_Command(message.From.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Данные указаны неверно, повторите попытку ещё раз.")
			return nil
		}

		if err := strings.Split(text[0], "@"); len(err) == 1 {
			Reset_User_Command(message.From.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Логин указан неверно, пожалуйста, повторите попытку.\nНапишите логин и пароль ещё раз.")
			return nil
		}

		login = _text[0]
		password = _text[1]
	
	} else {
		Reset_User_Command(message.From.ID, "reset_login")
		log.Printf("new command: reset_login")
		b.setMessage(message, "Данные указаны неверно, повторите попытку ещё раз.")
		return nil
	}

	//обнулим статус команды
	Delete_User_Command(message.From.ID)


	_, err := database.AuthenticateUser(login, password)
	if err == nil {
		database.AddUser(&models.User{Login: login, Password: password})
	} else {
		if err.Error() == "" {
			Reset_User_Command(message.From.ID, "reset_login")
			b.setMessage(message, "Неправильный проль, повторите попытку ещё раз.")
		}
		return nil
	}


	Push_Login_And_Password(message.From.ID, login, password)

	return nil
}


func (b *Bot) Auth(message *tgbotapi.Message) error {
	bot_user, err := Get_User_Command(message.From.ID)
	if err != nil {
		return err
	}


	user, err := database.AuthenticateUser(bot_user.Login, bot_user.Password)
	if err != nil {
		return err
	}


	//отправим сообщение в чат
	_err := b.setMessage(message, "Ваш логин: " + user.Login)
	if _err != nil {
		return _err
	}

	return nil
}















