package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Login    string             `json:"login" bson:"login"`
	Password string             `json:"password" bson:"password"`
	FIO      string             `json:"fio" bson:"fio"`
	User_ID  string             `json:"user_id" bson:"user_id"`
	RoleUser Role               `json:"roleuser" bson:"roleuser"`
	Logined  bool               `json:"logined" bson:"logined"`
	Tg_id    Id_telegram        `json:"tg_id" bson:"tg_id"`
}

type Role struct {
	Student             bool `json:"student" bson:"student"`
	Сurator             bool `json:"curator" bson:"curator"`
	Methodist           bool `json:"methodist" bson:"methodist"`
	SeniorMethodologist bool `josn:"seniormethodologist" bson:"seniormethodologist"`
	Director            bool `json:"director" bson:"director"`
}

type Group struct {
	Group string `json:"group" bson:"group"`
}

type Id_telegram struct {
	Id_telegram int64 `json:"id_telegram" bson:"id_telegram"`
}

type Passage struct {
	Passage_ID    primitive.ObjectID `bson:"_id"`
	FIO_student   string             `json:"fio_student" bson:"fio_student"`
	Group_student Group              `json:"group_student" bson:"group_student"`
	Passage_At    time.Time          `json:"passage_at" bson:"passage_at"`
	Flag          bool               `json:"flag" bson:"flag"`
}

type Subscription struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
}

type SubscriptionUser struct {
	ID             primitive.ObjectID `bson:"_id"`
	Tg_id          Id_telegram        `json:"id_tg" bson:"id_tg"`
	SubscriptionID primitive.ObjectID `bson:"subscription_id"`
	StartDate      time.Time          `bson:"start_date"`
	EndDate        time.Time          `bson:"end_date"`
	IsActive       bool               `bson:"is_active"`
}

type Order struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	SubscriptionUserID primitive.ObjectID `bson:"subscription_user_id"`
	Date               time.Time          `bson:"date"`
}

type Payment struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	OrderID    primitive.ObjectID `bson:"order_id"`
	Amount     float64            `bson:"amount"`
	PaymentURL string             `bson:"payment_url"`
}
