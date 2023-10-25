package main

import (
	"time"

	"github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/models"
)

func main() {
	// database.AddUser(&models.User{Login: "Owner", Password: "1234", FIO_student: "Mem"})
	// database.AddUserTelegram(3030)
	// database.AuthenticateUser(3030, "Owner", "1234")
	// log.Println(database.GetAllPassages())


	database.AddPassage(models.Passage{FIO_student: "Smolkin"})

	time.Sleep(time.Second * 10)

	database.AddPassage(models.Passage{FIO_student: "Admin"})

	time.Sleep(time.Second * 10)

	database.AddPassage(models.Passage{FIO_student: "Smolkin"})

	time.Sleep(time.Second * 10)

	database.AddPassage(models.Passage{FIO_student: "Admin"})

	time.Sleep(time.Second * 10)

	database.AddPassage(models.Passage{FIO_student: "Smolkin"})

	time.Sleep(time.Second * 10)

	database.AddPassage(models.Passage{FIO_student: "Admin"})
	// for {
	// 	user, passage, err := database.SearchItemInDB()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.Println(user, passage)

	// 	time.Sleep(10 * time.Second)
	// }

	// database.AddPassage(models.Passage{FIO_student: "Smolkin"})

}
