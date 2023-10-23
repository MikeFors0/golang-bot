package telegram

import (
	"fmt"
	"log"

	// "time"


	"github.com/MikeFors0/golang-bot/pkg/models"
	"github.com/MikeFors0/golang-bot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// запускает бота
func (b *Bot) Start() error {
	log.Printf("Authorized on account %v", &b.bot.Self)

	update, err := b.initUbdateChanel()
	if err != nil {
		return err
	}

	go b.GetPassage()

	b.handleUpdates(update)

	return nil
}

// авторизация, пользователь вводит свой логин и пароль
// может сделать это через пробел, или с переносом на следующую строку
func (b *Bot) Reg(message *tgbotapi.Message) error {

	log.Println("Значение в контексте пользователя при выполнении командв старт: " + User_comand[message.Chat.ID])

	login, password := b.handleLogin(message)
	if login == "" || password == "" {
		return nil
	}

	_, err := database.AuthenticateUser(database.Client, message.Chat.ID, login, password)
	if err != nil {
		if err.Error() == "invalid password" {
			Reset_User_Command(message.Chat.ID, "reset_login")
			b.setMessage(message.Chat.ID, "Неправильный проль, повторите попытку ещё раз.")
			return nil
		}
		// else {
		// 	err := database.AddUser(&models.User{Login: login, Password: password, Tg_id: models.Id_telegram{Id_telegram: message.Chat.ID}})
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	b.setMessage(message.Chat.ID, "Данные сохранены, чтобы проверить напишите /auth")

	Delete_User_Command(message.Chat.ID)

	log.Println("После выполнения команды старт, у пользователя установлена команда: " + User_comand[message.Chat.ID])

	return nil
}

func (b *Bot) Auth(message *tgbotapi.Message) error {

	log.Println("При вызове Auth команда подьзователя: " + User_comand[message.Chat.ID])

	user, err := database.GetUser(message.Chat.ID)
	if err != nil {
		if err.Error() == "user not found" {
			return b.setMessage(message.Chat.ID, "Нет такого пользователя, повторите попытку ещё раз, вызвав команду /start")
		}
		return b.setMessage(message.Chat.ID, err.Error())
	}

	//отправим сообщение в чат
	_err := b.setMessage(message.Chat.ID, "Ваш логин: "+user.Login)
	if _err != nil {
		return _err
	}

	return nil
}

func SendDataToUser(_bot *Bot, chatId int64, passage models.Passage) error {
	// создание нового сообщения
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("New data: %s", passage))

	// отправка сообщения
	_, err := _bot.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
