package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/models"
)

func main() {
	// group := models.Group{"П-21-01"}
	database.AddUser(&models.User{Login: "Owner", Password: "1234", FIO: "Смолькин Матвей Андреевич", RoleUser: models.Role{false, false, true, false, false}})
	database.AddUserTelegram(4355)
	database.AuthenticateUser(4355, "Owner", "1234")
	log.Println(database.GetUser(4355))

	// database.AddPassage(models.Passage{FIO_student: "Смолькин Матвей Андреевич", Group_student: models.Group{"ОБ-21-01"}})
	// database.AddPassage(models.Passage{FIO_student: "Ghhy", Group_student: models.Group{"ОБ-21-01"}})
	// database.AddPassage(models.Passage{FIO_student: "Ghhy", Group_student: models.Group{"П-21-01"}})

	// ps, us, _ := database.SearchItemInDB()
	// log.Println(ps, us)
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
