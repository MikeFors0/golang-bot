package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id"`
	Login              string             `json:"login" bson:"login"`
	Password           string             `json:"password" bson:"password"`
	FIO_student        string             `json:"fio_student" bson:"fio_student"`
	User_ID            string             `json:"user_id" bson:"user_id"`
	User_Subscriprtion Subscription_User  `json:"user_subscriprtion" bson:"user_subscriprtion"`
	Order_Status       []Order            `json:"order_status" bson:"order_status"`
	Passage_student    []Passage          `json:"passage_student" bson:"passage_student"`
	Logined            bool               `json:"logined" bson:"logined"`
	Tg_id              Id_telegram        `json:"tg_id" bson:"tg_id"`
}

type Subscription struct {
	Subscription_ID   primitive.ObjectID `bson:"_id"`
	Subscription_Name *string            `json:"subscription_name" bson:"subscription_name"`
	Price             *uint64            `json:"price" bson:"price"`
}

type Subscription_User struct {
	Subscription_ID   primitive.ObjectID `bson:"_id"`
	Subscription_Name *string            `json:"subscription_name" bson:"subscription_name"`
	Price             *uint64            `json:"price" bson:"price"`
}

type Order struct {
	Order_ID       primitive.ObjectID  `bson:"_id"`
	Order_Cart     []Subscription_User `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time           `json:"ordered_at" bson:"ordered_at"`
	Price          *uint64             `json:"price" bson:"price"`
	Payment_Method Payment             `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	digital bool
}

type Passage struct {
	Passage_ID  primitive.ObjectID `bson:"_id"`
	FIO_student string             `json:"fio_student" bson:"fio_student"`
	Passage_At  time.Time          `json:"passage_at" bson:"passage_at"`
}

type Id_telegram struct {
	Id_telegram int64 `json:"id_telegram" bson:"id_telegram"`
}
