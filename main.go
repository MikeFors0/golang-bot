package main

import (
	"log"

	"github.com/MikeFors0/golang-bot/database"
	"github.com/MikeFors0/golang-bot/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/MikeFors0/golang-bot/models"
)

func main() {
	user := models.User{
		Login:       "admin",
		FIO_student: "qwqw",
		Password:    "1234",
	}
	database.AddUser(&user)
	_, err := database.AddUserTelegram(database.Client, 1212)
	if err != nil {
		log.Panic(err)
	}

	_, err = database.AuthenticateUser(database.Client, 1212, "admin", "1234")
	if err != nil {
		log.Panic(err)
	}
	sb := models.Subscription{
		Name:        "rtghtrghrtthrhr",
		Description: "frfyyr",
		Price:       10,
	}
	database.AddSubscription(&sb)
	sbID, _ := primitive.ObjectIDFromHex("6530fc071fe2f8b719ccd448")
	database.BuySubscription(1212, sbID)
	database.CheckSubscription(1212)

	// passage := models.Passage{
	// 	FIO_student: "gnome",
	// }
	// database.AddPassage(passage)
	// database.CheckUserArrayVariable(1212, "[]Passage")
	// database.GetAllPassages()

	// database.CheckNewData()
	// database.GetUserByFIO("Smolkin")
	// time.Sleep(3 * time.Second)
	// rr := database.AddData(passage)
	// fmt.Println(rr)
	// for _, user := range users {
	// 	log.Println(user)
	// }
	// err := database.SendMessage()
	// if err != nil {
	// 	fmt.Println("error checking students")
	// 	// handle error
	// }
	// fmt.Println("Its Okey")

}
