package telegram

import (
	// "context"
	// "fmt"
	"fmt"
	"log"
	"strings"
	"sync"

	// "time"

	"github.com/MikeFors0/golang-bot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// обработчик сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message, wg *sync.WaitGroup) error {
	defer wg.Done()

	user, err := Get_User_Comand(message.Chat.ID)
	if err != nil {
		return err
	}

	switch user {
	case "start":
		return b.Reg(message)

	case "reset_login":
		return b.Reg(message)

	default:
		return b.setMessage(message, "К сожалению, я не знаю такой команды =(")
	}
}

// обработчик команд
func (b *Bot) handleCommand(chat_id int64, message *tgbotapi.Message, wg *sync.WaitGroup) error {
	defer wg.Done()

	switch message.Command() {
	case "start":
		return b.handleStart(message)

	case "auth":
		return b.Auth(message)
	case "buy":
		return b.buy(message)
	default:
		return b.setMessage(message, "К сожалению, я не знаю такой команды =((")
	}

	// return nil
}

func (b *Bot) handleStart(message *tgbotapi.Message) error {

	_, err := database.AddUserTelegram( message.Chat.ID)
	if err != nil {
		return err
	}

	log.Println("Добавили пользователя в бд")

	___err := b.setMessage(message, "Здравствуй, дорогой пользователь!\nДобро пожаловать в систему помощника по просмотру посещаемости учеников Самарского Государственного Колледжа.\nЯ буду отправлять Вам уведомления, когда Ваш ребёнок придёт в колледж.\nНапишите мне свои логин и пароль как на нашем сайте в любом из форматов ниже:\n\nuser@gmail.com 1234\n\nuser@gmail.com\n1234")
	if ___err != nil {
		return ___err
	}

	log.Println("Отправили сообщение")

	err = Set_User_Command(message.Chat.ID)
	if err != nil {
		return err
	}

	log.Println("Обратились к Set_User_Command")

	// log.Println("После handleStart у пользователя установлена команда: " + fmt.Sprint(ctx.Value(message.Chat.ID)))

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
			Reset_User_Command(message.Chat.ID, "reset_login")
			log.Printf("new command: reset_login")
			b.setMessage(message, "Данные указаны неверно, повторите попытку ещё раз.")
			return "", ""
		}

		login = _text[0]
		password = _text[1]

	} else {
		Reset_User_Command(message.Chat.ID, "reset_login")
		log.Printf("new command: reset_login")
		b.setMessage(message, "Данные указаны неверно, повторите попытку ещё раз.")
		return "", ""
	}

	return login, password
}

func (bot *Bot) HandlePreCheckoutQuery(update *tgbotapi.Update) {
	pca := tgbotapi.PreCheckoutConfig{
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
		OK:                 true,
	}

	_, err := bot.bot.AnswerPreCheckoutQuery(pca)
	if err != nil {
		log.Println("handlePreCheckount",err)
	}
}

func (bot *Bot) HandleSuccessfulPayment(update tgbotapi.Update) {
	message := update.Message

	paymentInfo := message.SuccessfulPayment

	paymentMessage := fmt.Sprintf("Платеж на сумму %d %s прошел успешно!!!",
		paymentInfo.TotalAmount/100, paymentInfo.Currency)

	msg := tgbotapi.NewMessage(message.Chat.ID, paymentMessage)
	_, err := bot.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

// func (b *Bot) handleRequest(message *tgbotapi.Message) {
// 	log.Println("Обработка запроса: " + message.Chat.UserName + " " + message.Text)
// 	time.Sleep(time.Second * 2)
// }
