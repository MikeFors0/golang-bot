package main

import (
	"log"
	"time"

	"github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/models"
)

func main() {
	database.AddUser(&models.User{Login: "Lox", Password: "1234", FIO: "Сергеевич Валентив"})
	database.AddUserTelegram(5050)
	database.AuthenticateUser(5050, "Lox", "1234")
	database.AddPassage(models.Passage{FIO_student: "Генадий Викторич Пета"})
	database.GetAllPassages()
	for {
		user, passage, err := database.SearchItemInDB()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(user, passage)
		time.Sleep(10 * time.Second)
	}

	// database.AddPassage(models.Passage{FIO_student: "Smolkin"})

}
