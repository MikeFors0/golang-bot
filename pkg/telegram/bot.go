package telegram

import (
	"fmt"
	"log"
	"strings"
	"time"

	// "github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

// проверка обновлений (нет ли новых сообщений)
func (b *Bot) initUbdateChanel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	update, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return update, nil
}

// отправить сообщение в чат
// принимает данные чата и текст сообщения, которое мы хотим отправить
func (b *Bot) setMessage(userId int64, text string) error {
	msg := tgbotapi.NewMessage(userId, text)
	_, err := b.bot.Send(msg)
	return err
}


func createMenu() tgbotapi.ReplyKeyboardMarkup {
	menu := tgbotapi.NewReplyKeyboard(
	 tgbotapi.NewKeyboardButtonRow(
	  tgbotapi.NewKeyboardButton("Мои данные"),
	  tgbotapi.NewKeyboardButton("Купить подписку"),
	 ),
	)
	return menu
   }


// отправка сообщений о посещении
func (b *Bot) GetPassage() {
	for {
		user, passage, err := database.SearchItemInDB()
		if err != nil {
			log.Println(err)
			continue
		}

		if user != nil {
			text_passage := strings.Split(fmt.Sprint(passage), " ")
			_time := strings.Split(text_passage[6], ".")
			b.setMessage(user.Tg_id.Id_telegram, "Ученик: "+text_passage[1] + " " + text_passage[2] + " " + text_passage[3] +"\nПрошёл через турникет в колледже \nДата: "+text_passage[5]+"\nВремя: "+ _time[0])
		}

		log.Println(user, passage)

		time.Sleep(5 * time.Second)
	}
}
