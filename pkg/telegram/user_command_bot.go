package telegram

import (
	"errors"
	"fmt"
)

type User struct {
	ID int64
	Login string
	Password string
	Command string
}

var user_data = map[int64]User{}

func Set_User_Command(userId int64) error {

	if user_data[userId].ID != 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	user_data = make(map[int64]User)

	user_data[userId] = User{
		ID: userId,
		Command: "start",
	}

	return nil
}

func Get_User_Command(userId int64) (*User, error) {
	if userId == 0 {
		return nil, errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	user := user_data[userId]

	return &user, nil
}

func Reset_User_Command(userId int64, new_command string) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	user_data[userId] = User{
		Command: new_command,
	}

	return nil
}

func Push_Login_And_Password(userId int64, login, password string) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	user_data[userId] = User{
		Login: login,
		Password: password,
	}

	return nil
}


func Delete_User_Command(userId int64) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	user_data[userId] = User{
		Command: "none",
	}

	return nil

}
