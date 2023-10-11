package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MikeFors0/golang-bot/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = UserData(Client, "users")

var validate = validator.New()

func AddUser(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	count, err := UserCollection.CountDocuments(ctx, bson.M{"login": user.Login})
	defer cancel()
	if err != nil {
		log.Panic(err)
		fmt.Println("Error login, checking login")

	}

	if count > 0 {
		defer cancel()
		fmt.Println("user already exists")
	}else {
		validatorErr := validate.Struct(user)
		defer cancel()
		if validatorErr != nil {
			log.Panic(validatorErr.Error())
			fmt.Println(validatorErr.Error())
		}
	
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			fmt.Sprintf("User item was not created: %v", insertErr)
		}
		defer cancel()
		fmt.Println("User created")
		return nil
	}
	return nil

}

func GetUsers() (*[]models.User, error) {
	cur, err := UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User
	for cur.Next(context.Background()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return &users, nil
}
