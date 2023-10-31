package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/pkg/database"
)

func main() {
	database.AddUserTelegram(3030)
	database.AuthenticateUser(3030, "Owner", "1234")
	log.Println(database.GetAllPassages())
	// database.AddPassage(models.Passage{FIO_student: "Mem"})
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
