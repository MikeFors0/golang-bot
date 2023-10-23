package telegram

import (
	"errors"
	"fmt"
	"log"
)

var User_comand map[int64]string


func Set_User_Command(userId int64) error {

	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}


	User_comand[userId] = "start"

	log.Println("Пользователю с id: " + fmt.Sprint(userId) + " присовоена новая команда -> " + User_comand[userId])

	return nil
}

func Get_User_Comand(userId int64) (string, error) {

	if userId == 0 {
		return "", errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	return User_comand[userId], nil
}

func Reset_User_Command(userId int64, new_command string) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	log.Println("Команда пользователя до изменений: " + User_comand[userId])

	User_comand[userId] = new_command

	log.Println("Пользователю с id: " + fmt.Sprint(userId) + " присовоена новая команда -> " + 	User_comand[userId])

	return nil
}


func Delete_User_Command(userId int64) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	User_comand[userId] = "nil"

	log.Println("Пользователю с id: " + fmt.Sprint(userId) + " присовоена новая команда -> " + User_comand[userId])

	return nil
}
