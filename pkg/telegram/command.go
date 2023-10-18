package telegram

import (
	"fmt"
	"log"
	"time"

	"github.com/MikeFors0/golang-bot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// запускает бота
func (b *Bot) Start() error {
	log.Printf("Authorized on account %v", &b.bot.Self)

	update, err := b.initUbdateChanel()
	if err != nil {
		return err
	}

	b.handleUpdates(update)

	return nil
}

// авторизация, пользователь вводит свой логин и пароль
// может сделать это через пробел, или с переносом на следующую строку
func (b *Bot) Reg(message *tgbotapi.Message) error {

	login, password := b.handleLogin(message)
	if login == "" || password == "" {
		return nil
	}


	_, err := database.AuthenticateUser(database.Client, message.Chat.ID, login, password)
	if err != nil {
		if err.Error() == "invalid password" {
			Reset_User_Command(user_comand_context, message.Chat.ID, "reset_login")
			b.setMessage(message, "Неправильный проль, повторите попытку ещё раз.")
			return nil
		}
		// else {
		// 	err := database.AddUser(&models.User{Login: login, Password: password, Tg_id: models.Id_telegram{Id_telegram: message.Chat.ID}})
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	b.setMessage(message, "Данные сохранены, чтобы проверить напишите /auth")

	Reset_User_Command(user_comand_context, message.Chat.ID, "nil")


	//после проверки убрать следующие строчки
	user, __err := Get_User_Comand(user_comand_context, message.Chat.ID)
	if __err != nil {
		return __err
	}

	log.Println("После выполнения команды старт, у пользователя установлена команда: " + fmt.Sprint(user))

	time.Sleep(time.Second * 3)

	return nil
}

func (b *Bot) Auth(message *tgbotapi.Message) error {

	user, err := database.GetUser(message.Chat.ID)
	if err != nil {
		return b.setMessage(message, err.Error())
	}


	//отправим сообщение в чат
	_err := b.setMessage(message, "Ваш логин: " + user.Login)
	if _err != nil {
		return _err
	}

	return nil
}
