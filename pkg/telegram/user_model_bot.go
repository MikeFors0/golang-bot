package telegram

import (
	"fmt"
	"log"
)

type User struct {
	Id           int
	Runs_Command string
}

var map_User = map[int]User{}

func SetUser(userId int) error {

	if map_User[userId].Id != 0 {
		return fmt.Errorf("this user already exists %q", userId)
	}

	map_User := make(map[int]User)

	map_User[userId] = User{
		userId,
		"",
	}

	log.Println(map_User[userId])

	return nil
}

func GetUser(userId int) (*User, error) {
	if map_User[userId].Id == 0 {
		return nil, fmt.Errorf("invalid user id %q", userId)
	}

	user := map_User[userId]
	return &user, nil
}

func Reset_Runs_Command(userId int, new_command string) error {
	user, err := GetUser(userId)
	if err != nil {
		return err
	}

	user.Runs_Command = new_command

	log.Printf("user command: %v", user)

	return nil
}
