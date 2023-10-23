package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Login           string             `json:"login" bson:"login"`
	Password        string             `json:"password" bson:"password"`
	FIO_student     string             `json:"fio_student" bson:"fio_student"`
	User_ID         string             `json:"user_id" bson:"user_id"`
	Passage_student []Passage          `json:"passage_student" bson:"passage_student"`
	Logined         bool               `json:"logined" bson:"logined"`
	Tg_id           Id_telegram        `json:"tg_id" bson:"tg_id"`
}

type Subscription struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
}

type SubscriptionUser struct {
	ID             primitive.ObjectID `bson:"_id"`
	Tg_id          Id_telegram        `bson:"tg_id"`
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

type Passage struct {
	Passage_ID  primitive.ObjectID `bson:"_id"`
	FIO_student string             `json:"fio_student" bson:"fio_student"`
	Passage_At  time.Time          `json:"passage_at" bson:"passage_at"`
}

type Id_telegram struct {
	Id_telegram int64 `json:"id_telegram" bson:"id_telegram"`
}
