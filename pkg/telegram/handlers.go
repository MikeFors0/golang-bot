package telegram

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/MikeFors0/golang-bot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// обработчик обновлений
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	var wg sync.WaitGroup

	for update := range updates {

		log.Println("\n\nПолучено новое сообщение: " + "\nusername -> " + update.Message.Chat.UserName + "\ntext -> " + update.Message.Text + "\nchat_id -> " + fmt.Sprint(update.Message.Chat.ID) + "\nКонец\n")

		if update.PreCheckoutQuery != nil {

			b.HandlePreCheckoutQuery(&update)

		} else if update.Message.SuccessfulPayment != nil {

			b.HandleSuccessfulPayment(update.Message)

		}

		//если обновлений нет, продолжит ожидать
		if update.Message == nil {
			continue
		}

		//если это команда, перейдём в обработчик команд
		if update.Message.IsCommand() {
			wg.Add(1)
			go b.handleCommand(update.Message.Chat.ID, update.Message, &wg)
			time.Sleep(time.Second * 1)
			continue
		} else { //если текст, перейдём в обработчик сообщений
			wg.Add(1)
			go b.handleMessage(update.Message, &wg)
			time.Sleep(time.Second * 1)
			continue
		}
	}

	wg.Wait()
}

// обработчик сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message, wg *sync.WaitGroup) error {
	defer wg.Done()

	user, err := Get_User_Comand(message.Chat.ID)
	if err != nil {
		return err
	}

	switch message.Text {
		case "Мои данные":
			return b.Auth(message)
		case "Купить подписку":
			return b.buy(message)
		
		default:
			switch user {
				case "start":
					return b.Reg(message)

				case "reset_login":
					return b.Reg(message)

				default:
					return b.setMessage(message.Chat.ID, "К сожалению, я не знаю такой команды =(")
				}
	}

}

// обработчик команд
func (b *Bot) handleCommand(chat_id int64, message *tgbotapi.Message, wg *sync.WaitGroup) error {
	defer wg.Done()

	switch message.Command() {
	case "start":
		return b.handleStart(message)

	default:
		return b.setMessage(message.Chat.ID, "К сожалению, я не знаю такой команды =((")
	}

}

// действия при вызове /start
func (b *Bot) handleStart(message *tgbotapi.Message) error {

	_, err := database.AddUserTelegram(message.Chat.ID)
	if err != nil {

		if err.Error() == "пользователь с таким ID уже существует" {
			b.setMessage(message.Chat.ID, "С возвращением, дорогой пользователь!\nЯ вас помню 🤗")
			return nil
		}

		return err
	}


	err = Set_User_Command(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуй, дорогой пользователь!\nДобро пожаловать в систему помощника по просмотру посещаемости учеников Самарского Государственного Колледжа.\nЯ буду отправлять Вам уведомления, когда Ваш ребёнок придёт в колледж.\nНапишите мне свои логин и пароль как на нашем сайте в любом из форматов ниже:\n\nAdmin 4444\n\nAdmin\n4444")
	msg.ReplyMarkup = createMenu()
	b.bot.Send(msg)

	return nil
}

// вспомогательная функция обработки введённых пользоватлем данных
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
			Reset_User_Command(message.Chat.ID, "reset_login")
			b.setMessage(message.Chat.ID, "Данные указаны неверно, повторите попытку ещё раз.")
			return "", ""
		}

		login = _text[0]
		password = _text[1]

	} else {
		Reset_User_Command(message.Chat.ID, "reset_login")
		b.setMessage(message.Chat.ID, "Данные указаны неверно, повторите попытку ещё раз.")
		return "", ""
	}

	return login, password
}

// обработчики покупки
func (b *Bot) HandlePreCheckoutQuery(update *tgbotapi.Update) tgbotapi.PreCheckoutConfig {
	pca := tgbotapi.PreCheckoutConfig{
		OK:                 true,
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
	}
	_, err := b.bot.AnswerPreCheckoutQuery(pca)

	if err != nil {
		log.Println("handlePreCheckount", err)
	}
	log.Println(pca)
	return pca
}

func (b *Bot) HandleSuccessfulPayment(message *tgbotapi.Message) *tgbotapi.SuccessfulPayment {
	paymentInfo := message.SuccessfulPayment

	paymentMessage := fmt.Sprintf(
		"Платеж на сумму %d %s прошел успешно!!!",
		paymentInfo.TotalAmount/100,
		paymentInfo.Currency,
	)
	log.Println(paymentInfo)
	b.setMessage(message.Chat.ID, paymentMessage)
	return paymentInfo

}
