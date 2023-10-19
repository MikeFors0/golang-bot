package telegram

import (
	"context"
	"errors"
	"fmt"
	"log"
)


var user_comand_context context.Context

func Init_Context() {
	user_comand_context = context.Background()
}


func Set_User_Command(ctx context.Context, userId int64) error {

	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	_ = context.WithValue(ctx, fmt.Sprint(userId), "start")

	log.Println("Пользователю с id: " + fmt.Sprint(userId) + " присовоена новая команда -> " + fmt.Sprint(ctx.Value(userId)))

	return nil
}

func Get_User_Comand(ctx context.Context, userId int64) (any, error) {

	if userId == 0 {
		return nil, errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	return ctx.Value(userId), nil
}

func Reset_User_Command(ctx context.Context, userId int64, new_command string) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	context.WithValue(ctx, userId, new_command)

	log.Println("Пользователю с id: " + fmt.Sprint(userId) + " присовоена новая команда -> " + fmt.Sprint(ctx.Value(userId)))

	return nil
}


func Delete_User_Command(ctx context.Context, userId int64) error {
	if userId == 0 {
		return errors.New("invalid user id:" + fmt.Sprintf("%c", userId))
	}

	_ = context.WithValue(ctx, userId, "nil")

	log.Println("Пользователю с id: " + fmt.Sprint(userId) + " присовоена новая команда -> " + fmt.Sprint(ctx.Value(userId)))

	return nil
}
