package telegram

import (
	"log"

	"github.com/MikeFors0/golang-bot/database"
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
	Delete_User_Command(message.From.ID)


	_, err := database.AuthenticateUser(database.Client, message.From.ID, login, password)
	if err != nil {
		if err.Error() == "invalid password" {
			Reset_User_Command(message.From.ID, "reset_login")
			b.setMessage(message, "Неправильный проль, повторите попытку ещё раз.")
			return nil
		}
		return err
	}


	Push_Login_And_Password(message.From.ID, login, password)

	return nil
}


func (b *Bot) Auth(message *tgbotapi.Message) error {
	bot_user, err := Get_User_Command(message.From.ID)
	if err != nil {
		return err
	}


	user, err := database.AuthenticateUser(database.Client, bot_user.ID, bot_user.Login, bot_user.Password)
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















