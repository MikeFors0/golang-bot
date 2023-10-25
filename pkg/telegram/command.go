package telegram

import (
	"log"

	"github.com/MikeFors0/golang-bot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var PAYMENTS_TOKEN = "401643678:TEST:12f8fa5c-82aa-47e6-b808-2f1fd95ad954"

var PRICE = tgbotapi.LabeledPrice{
	Label:  "Подписка на 1 месяц",
	Amount: 500 * 100,
}

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

	_, err := database.AuthenticateUser(message.Chat.ID, login, password)
	if err != nil {
		if err.Error() == "invalid password" {
			Reset_User_Command(message.Chat.ID, "reset_login")
			b.setMessage(message.Chat.ID, "Неправильный проль, повторите попытку ещё раз.")
			return nil
		} else {
			b.setMessage(message.Chat.ID, err.Error())
		}
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
		if err.Error() == "пользователь не найден" {
			return b.setMessage(message.Chat.ID, "Нет такого пользователя, повторите попытку ещё раз, вызвав команду /start")
		}
		return b.setMessage(message.Chat.ID, err.Error())
	}

	//отправим сообщение в чат
	_err := b.setMessage(message.Chat.ID, "Ваш логин: " + user.Login)
	if _err != nil {
		return _err
	}

	return nil
}















func (b *Bot) buy(message *tgbotapi.Message) error {

	invoice := tgbotapi.NewInvoice(
		message.Chat.ID,
		"Подписка на бота",
		"Активация подписки на бота на 1 месяц",
		"test",
		PAYMENTS_TOKEN,
		"one-month-subscription",
		"RUB",
		&[]tgbotapi.LabeledPrice{PRICE},
	) 

	invoice.PhotoURL = "https://i.ytimg.com/vi/ntoyQN_0sMY/maxresdefault.jpg"
	invoice.PhotoWidth = 416
	invoice.PhotoHeight = 234
	invoice.PhotoSize = 416


	_, _err := b.bot.Send(invoice)
	if _err != nil {
		log.Println("invoice err", _err)
		return  _err
	}
	
	return nil
}
