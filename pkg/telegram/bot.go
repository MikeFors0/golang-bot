package telegram

import (
	"fmt"
	"log"
	"strings"
	"sync"
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

// обработчик обновлений
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	var wg sync.WaitGroup

	for update := range updates {

		log.Println("\n\nПолучено новое сообщение: " + "\nusername -> " + update.Message.Chat.UserName + "\ntext -> " + update.Message.Text + "\nchat_id -> " + fmt.Sprint(update.Message.Chat.ID) + "\nКонец\n")

		

		if update.PreCheckoutQuery != nil {
			log.Println("err1")
			b.HandlePreCheckoutQuery(&update)
			return
		  } else if update.Message != nil && update.Message.SuccessfulPayment != nil {
			log.Println("err2")
			b.HandleSuccessfulPayment(update.Message)
			log.Println(update.Message)
			log.Println(&update)
			return
		  }
		log.Println(update.Message)
		log.Println(&update)

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

		//если обновлений нет, продолжит ожидать
		if update.Message == nil {
			continue
		}

	}
	log.Println(updates)

	wg.Wait()
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

func (b *Bot) GetPassage() {

	for {
		user, passage, err := database.SearchItemInDB()
		if err != nil {
			log.Println(err)
			continue
		}

		if user != nil {
			text_passage := strings.Split(fmt.Sprint(passage), " ")
			b.setMessage(user.Tg_id.Id_telegram, "Ученик с фамилией: " + text_passage[1] + "\nПрошёл через турникет в колледже \nДата: " + text_passage[2] + "\nВремя: " + text_passage[3])
		}

		log.Println(user, passage)

		time.Sleep(5 * time.Second)
	}
}
