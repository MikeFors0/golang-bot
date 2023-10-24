package main

import (
	"github.com/MikeFors0/golang-bot/pkg/database"
	"github.com/MikeFors0/golang-bot/pkg/models"
)

func main() {
	group := models.Group{"ОБ-21-01"}
	group2 := models.Group{"ПО-21-01"}
	database.AddUser(&models.User{Login: "Curator", Password: "1234", FIO_student: "Curator", GroupUser: []models.Group{group, group2}, RoleUser: models.Role{false, true, false, false, false}})
	database.AddUserTelegram(4040)
	database.AuthenticateUser(4040, "Curator", "1234")
	database.GetUser(4040)

	// database.AddPassage(models.Passage{FIO_student: "Mem", Group_student: models.Group{"ОБ-21-01"}})
	// database.AddPassage(models.Passage{FIO_student: "Lox", Group_student: models.Group{"ОБ-21-01"}})
	// database.AddPassage(models.Passage{FIO_student: "Ghhy", Group_student: models.Group{"ОБ-21-01"}})
	// database.AddPassage(models.Passage{FIO_student: "Ghhy", Group_student: models.Group{"ПО-21-01"}})

	// res, _ := database.SearchForKurator(3030, group)
	// log.Println(res)
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
