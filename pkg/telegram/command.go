package telegram

import (
	
	"log"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//запускает бота
func (b *Bot) Start() error {
	log.Printf("Authorized on account %v", &b.bot.Self)

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
	login, password := b.handleLogin(message)
	if login == "" || password == "" {
		return nil
	}

	//обнулим статус команды
	Delete_User_Command(int(message.Chat.ID))


	_, err := database.AuthenticateUser(database.Client, uint(message.Chat.ID), login, password)
	if err != nil {
		if err.Error() == "invalid password" {
			Reset_User_Command(int(message.Chat.ID), "reset_login")
			b.setMessage(message, "Неправильный проль, повторите попытку ещё раз.")
			return nil
		} else {
			err := database.AddUser(&models.User{Login: login, Password: password, Tg_id: models.Id_telegram{Id_telegram: uint(message.Chat.ID)}})
			if err != nil {
				return err
			}
		} 
	}


	Push_Login_And_Password(int(message.Chat.ID), login, password)

	b.setMessage(message, "Данные сохранены, чтобы проверить напишите /auth")

	return nil
}


func (b *Bot) Auth(message *tgbotapi.Message) error {
	bot_user, err := Get_User_Command(int(message.Chat.ID))
	if err != nil {
		return err
	}


	user, err := database.AuthenticateUser(database.Client, uint(bot_user.ID), bot_user.Login, bot_user.Password)
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















